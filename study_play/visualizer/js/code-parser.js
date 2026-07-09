/**
 * Parses user code, detects data-structure patterns, and auto-instruments
 * with visualization calls when manual viz.step() is absent.
 */
const CodeParser = (function () {
  const POINTER_NAMES = new Set(["left", "right", "i", "j", "k", "lo", "hi", "mid", "start", "end", "slow", "fast"]);

  const TREE_SIGNALS = [/buildTreeFromLevelOrder/i, /\.left\b/, /\.right\b/, /TreeNode/, /inorder|preorder|postorder/i, /maxDepth|levelOrder/i];
  const GRAPH_SIGNALS = [/parseGraphInput/i, /grid\s*\[/, /image\s*\[/, /numIslands|floodFill|shortestPath/i];
  const STRING_SIGNALS = [/parseStringInput/i, /viz\.setString/, /\.charAt\s*\(/, /palindrome|anagram|substring/i, /s\[\s*\w+\s*\]/];
  const ARRAY_SIGNALS = [/viz\.setArray/, /parseArrayInput/, /\[\s*\.\.\.\s*data\s*\]/, /\w+\[\s*\w+\s*\]/];

  function parseAst(code) {
    if (typeof acorn === "undefined") throw new Error("Acorn parser not loaded");
    return acorn.parse(code, { ecmaVersion: 2022, locations: true, sourceType: "script", allowAwaitOutsideFunction: true });
  }

  function walk(node, visitor) {
    if (!node || typeof node !== "object") return;
    visitor(node);
    for (const key of Object.keys(node)) {
      const child = node[key];
      if (Array.isArray(child)) child.forEach((c) => c && c.type && walk(c, visitor));
      else if (child && child.type) walk(child, visitor);
    }
  }

  function getSource(code, node) {
    if (!node?.loc) return "";
    const lines = code.split("\n");
    return lines.slice(node.loc.start.line - 1, node.loc.end.line).join("\n");
  }

  function getLineIndent(code, lineNum) {
    const m = (code.split("\n")[lineNum - 1] || "").match(/^(\s*)/);
    return m ? m[1] : "  ";
  }

  function scoreByPatterns(code) {
    const scores = { array: 1, string: 0, tree: 0, graph: 0 };
    for (const re of TREE_SIGNALS) if (re.test(code)) scores.tree += 2;
    for (const re of GRAPH_SIGNALS) if (re.test(code)) scores.graph += 2;
    for (const re of STRING_SIGNALS) if (re.test(code)) scores.string += 2;
    for (const re of ARRAY_SIGNALS) if (re.test(code)) scores.array += 1;
    return scores;
  }

  function detectStructureType(code, ast) {
    const scores = scoreByPatterns(code);
    walk(ast, (node) => {
      if (node.type === "MemberExpression" && node.computed) {
        const obj = getSource(code, node.object);
        if (/grid|image|board|matrix/i.test(obj)) scores.graph += 3;
        else if (/^(s|str|chars|string)/i.test(obj.trim())) scores.string += 2;
        else scores.array += 1;
      }
      if (node.type === "MemberExpression" && !node.computed) {
        const prop = node.property.name || "";
        if (prop === "left" || prop === "right") scores.tree += 2;
      }
      if (node.type === "CallExpression") {
        const callee = getSource(code, node.callee);
        if (/buildTreeFromLevelOrder/.test(callee)) scores.tree += 5;
        if (/parseGraphInput/.test(callee)) scores.graph += 5;
        if (/parseStringInput/.test(callee)) scores.string += 5;
      }
    });
    const ranked = Object.entries(scores).sort((a, b) => b[1] - a[1]);
    const [type, score] = ranked[0];
    const second = ranked[1][1];
    return { structureType: type, confidence: score > 0 ? Math.min(1, score / (score + second + 1)) : 0.25, scores };
  }

  function detectPatterns(code, ast) {
    const patterns = new Set();
    walk(ast, (node) => {
      if (node.type === "ForStatement") patterns.add("loop");
      if (node.type === "WhileStatement") patterns.add("while-loop");
      if (node.type === "IfStatement" && /\w+\[\s*\w+\s*\]/.test(getSource(code, node.test))) patterns.add("compare");
    });
    for (const name of POINTER_NAMES) {
      if (new RegExp(`(?:let|const|var)\\s+${name}\\b`).test(code)) {
        patterns.add("pointers");
        if (name === "left" || name === "right") patterns.add("two-pointers");
        if (name === "mid") patterns.add("binary-search");
      }
    }
    if (/viz\.step|await\s+viz\./.test(code)) patterns.add("manual-viz");
    if (/queue|bfs|dfs/i.test(code)) patterns.add("graph-traversal");
    return [...patterns];
  }

  function findPrimaryDataVar(code, ast, structureType) {
    const candidates = [];
    walk(ast, (node) => {
      if (node.type !== "VariableDeclarator" || !node.id?.name || !node.init) return;
      const initSrc = getSource(code, node.init);
      if (/\[\s*\.\.\.\s*data\s*\]/.test(initSrc) || initSrc.trim() === "data") candidates.push({ name: node.id.name, priority: 10, line: node.loc.start.line, endLine: node.loc.end.line });
      else if (/String\s*\(\s*data/.test(initSrc)) candidates.push({ name: node.id.name, priority: 9, line: node.loc.start.line, endLine: node.loc.end.line });
      else if (/buildTreeFromLevelOrder/.test(initSrc)) candidates.push({ name: node.id.name, priority: 10, line: node.loc.start.line, endLine: node.loc.end.line });
      else if (/parseGraphInput|\.map\s*\(/.test(initSrc)) candidates.push({ name: node.id.name, priority: 8, line: node.loc.start.line, endLine: node.loc.end.line });
    });
    const defaults = { array: "arr", string: "s", tree: "root", graph: "grid" };
    if (!candidates.length) {
      for (const name of ["arr", "nums", "array", "s", "str", "grid", "image", "root"]) {
        if (new RegExp(`(?:const|let|var)\\s+${name}\\b`).test(code)) candidates.push({ name, priority: 5, line: 1 });
      }
    }
    candidates.sort((a, b) => b.priority - a.priority);
    return candidates[0] || { name: defaults[structureType] || "arr", line: 1, endLine: 1 };
  }

  function findRunFunction(ast) {
    let runFn = null;
    walk(ast, (node) => {
      if (node.type === "FunctionDeclaration" && node.id?.name === "run") runFn = node;
    });
    return runFn;
  }

  function hasManualViz(code) {
    return /viz\.step\s*\(|await\s+viz\.(step|setArray|setString|highlight)/.test(code);
  }

  function analyze(code) {
    try {
      const ast = parseAst(code);
      const { structureType, confidence, scores } = detectStructureType(code, ast);
      const patterns = detectPatterns(code, ast);
      const primary = findPrimaryDataVar(code, ast, structureType);
      return {
        ok: true,
        ast,
        structureType,
        confidence,
        scores,
        patterns,
        primaryVar: primary.name,
        primaryVarLine: (primary.endLine || primary.line) + 1,
        hasManualViz: hasManualViz(code),
        hasRunFunction: !!findRunFunction(ast),
      };
    } catch (err) {
      return {
        ok: false,
        error: err.message,
        structureType: "array",
        confidence: 0,
        patterns: [],
        primaryVar: "arr",
        primaryVarLine: 1,
        hasManualViz: hasManualViz(code),
        hasRunFunction: /function\s+run\b/.test(code),
      };
    }
  }

  function traceSnippet(indent, line, event, opts = {}) {
    const parts = [`line: ${line}`, `event: '${event}'`];
    if (opts.varNames?.length) parts.push(`vars: { ${opts.varNames.join(", ")} }`);
    if (opts.dataVar) parts.push(`data: ${opts.dataVar}`);
    if (opts.indices?.length) parts.push(`indices: [${opts.indices.join(", ")}]`);
    if (opts.highlightType) parts.push(`highlightType: '${opts.highlightType}'`);
    if (opts.message != null) parts.push(`message: ${JSON.stringify(String(opts.message))}`);
    return `${indent}await __autoViz.trace({ ${parts.join(", ")} });\n`;
  }

  function collectLoopVars(node, code) {
    const names = [];
    if (node.type === "ForStatement" && node.init?.type === "VariableDeclaration") {
      for (const d of node.init.declarations) {
        if (d.id?.name && POINTER_NAMES.has(d.id.name)) names.push(d.id.name);
      }
    }
    if (node.type === "WhileStatement") {
      const testSrc = getSource(code, node.test);
      for (const name of POINTER_NAMES) {
        if (new RegExp(`\\b${name}\\b`).test(testSrc)) names.push(name);
      }
    }
    return names;
  }

  function instrument(code, analysis) {
    if (!analysis.ok || analysis.hasManualViz) {
      return { code, instrumented: false, reason: "manual-viz" };
    }
    const ast = analysis.ast || parseAst(code);
    const runFn = findRunFunction(ast);
    if (!runFn?.body?.body?.length) {
      return { code, instrumented: false, reason: "no-run-function" };
    }

    const primary = analysis.primaryVar;
    const insertions = [];

    insertions.push({
      line: analysis.primaryVarLine,
      text: traceSnippet(getLineIndent(code, analysis.primaryVarLine), analysis.primaryVarLine, "start", {
        dataVar: primary,
        message: "Initialize visualization",
      }),
      priority: 0,
    });

    walk(runFn.body, (node) => {
      if ((node.type === "ForStatement" || node.type === "WhileStatement") && node.body?.type === "BlockStatement" && node.body.body.length) {
        const line = node.body.body[0].loc.start.line;
        const indent = getLineIndent(code, line) + "  ";
        const varNames = collectLoopVars(node, code);
        insertions.push({
          line,
          text: traceSnippet(indent, line, "loop", { varNames, dataVar: primary, message: "Loop iteration" }),
          priority: 1,
        });
      }

      if (node.type === "IfStatement") {
        const condSrc = getSource(code, node.test);
        const memberMatches = [...condSrc.matchAll(/(\w+)\[\s*(\w+)\s*\]/g)];
        if (memberMatches.length) {
          const line = node.loc.start.line;
          const indent = getLineIndent(code, line);
          const indexVars = [...new Set(memberMatches.map((m) => m[2]).filter((v) => /^\w+$/.test(v)))];
          insertions.push({
            line,
            text: traceSnippet(indent, line, "compare", {
              varNames: indexVars,
              dataVar: primary,
              indices: indexVars,
              highlightType: "compare",
              message: condSrc.replace(/`/g, "'").slice(0, 60),
            }),
            priority: 2,
          });
        }
      }

      if (node.type === "ExpressionStatement" && node.expression?.type === "AssignmentExpression" && node.expression.left?.type === "ArrayExpression") {
        const line = node.loc.start.line;
        const indent = getLineIndent(code, line);
        insertions.push({
          line: node.loc.end.line + 1,
          text: traceSnippet(indent, node.loc.end.line + 1, "swap", { dataVar: primary, highlightType: "swap", message: "Swap elements" }),
          priority: 3,
        });
      }
    });

    const byLine = new Map();
    for (const ins of insertions) {
      if (!byLine.has(ins.line)) byLine.set(ins.line, []);
      byLine.get(ins.line).push(ins);
    }

    const lines = code.split("\n");
    for (const lineNum of [...byLine.keys()].sort((a, b) => b - a)) {
      const combined = byLine.get(lineNum).sort((a, b) => a.priority - b.priority).map((i) => i.text).join("");
      lines[lineNum - 1] = combined + (lines[lineNum - 1] || "");
    }

    return { code: lines.join("\n"), instrumented: true, reason: "auto-instrumented", insertionCount: insertions.length };
  }

  function prepare(code, options = {}) {
    const analysis = analyze(code);
    let working = code;
    let instrumented = { code: working, instrumented: false, reason: "skipped" };

    if (options.autoInstrument !== false && !analysis.hasManualViz && analysis.hasRunFunction) {
      instrumented = instrument(working, analysis);
      if (instrumented.instrumented) working = instrumented.code;
    }

    return { analysis, code: working, instrumented: instrumented.instrumented, instrumentReason: instrumented.reason };
  }

  return { analyze, instrument, prepare };
})();

function createAutoViz(viz, config) {
  const { structureType } = config;

  function syncData(data) {
    if (data === undefined || data === null) return;
    if (Array.isArray(data)) {
      if (Array.isArray(data[0])) viz.setGraph(data.map((r) => [...r]));
      else viz.setArray([...data]);
    } else if (typeof data === "string") {
      viz.setString(data);
    } else if (data && typeof data === "object" && "val" in data) {
      viz.setTree(data);
    }
  }

  function resolveHighlightIndices(vars, indices) {
    const out = [];
    for (const key of indices) {
      if (typeof vars[key] === "number") out.push(vars[key]);
    }
    if (!out.length) {
      for (const [k, v] of Object.entries(vars)) {
        if (typeof v === "number") out.push(v);
      }
    }
    return [...new Set(out)];
  }

  return {
    async trace({ line, event, vars = {}, data, indices = [], highlightType = "compare", message = "" }) {
      syncData(data);

      const ptr = {};
      for (const [k, v] of Object.entries(vars)) {
        if (typeof v === "number") ptr[k] = v;
      }
      if (Object.keys(ptr).length) viz.setPointers(ptr);

      const idx = resolveHighlightIndices(vars, indices);
      if (idx.length) viz.highlight(idx, highlightType || event);

      await viz.step({
        line,
        message: message || `Auto: ${event}`,
        variables: vars,
      });
    },
  };
}

if (typeof module !== "undefined") {
  module.exports = { CodeParser, createAutoViz };
}
