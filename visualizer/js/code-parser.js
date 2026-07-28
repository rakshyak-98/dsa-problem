/**
 * Parses user code, detects data-structure patterns, and auto-instruments
 * with visualization calls when manual viz.step() is absent.
 */
const CodeParser = (function () {
  const POINTER_NAMES = new Set(["left", "right", "i", "j", "k", "lo", "hi", "mid", "start", "end", "slow", "fast", "r", "c", "nr", "nc", "row", "col"]);

  const TREE_SIGNALS = [/buildTreeFromLevelOrder/i, /\.left\b/, /\.right\b/, /TreeNode/, /inorder|preorder|postorder/i, /maxDepth|levelOrder/i, /stack\.push\s*\(\s*\w+\s*\)/];
  const GRAPH_SIGNALS = [/parseGraphInput/i, /\w+\[\s*\w+\s*\]\[\s*\w+\s*\]/, /numIslands|floodFill|shortestPath/i, /queue\.push\s*\(\s*\[/];
  const STRING_SIGNALS = [/parseStringInput/i, /viz\.setString/, /String\s*\(\s*data/, /\.charAt\s*\(/, /palindrome|anagram|substring/i];
  const ARRAY_SIGNALS = [/viz\.setArray/, /parseArrayInput/, /\[\s*\.\.\.\s*data\s*\]/];

  function parseAst(code) {
    if (typeof acorn === "undefined") throw new Error("Acorn parser not loaded");
    return acorn.parse(code, { ecmaVersion: 2022, locations: true, sourceType: "script", allowAwaitOutsideFunction: true });
  }

  function walk(node, visitor, parents = []) {
    if (!node || typeof node !== "object") return;
    const insideNestedFn = parents.some((p) =>
      ["FunctionDeclaration", "FunctionExpression", "ArrowFunctionExpression"].includes(p.type)
    );
    visitor(node, insideNestedFn);
    const next = [...parents, node];
    for (const key of Object.keys(node)) {
      const child = node[key];
      if (Array.isArray(child)) child.forEach((c) => c && c.type && walk(c, visitor, next));
      else if (child && child.type) walk(child, visitor, next);
    }
  }

  function getSource(code, node) {
    if (!node?.loc) return "";
    const lines = code.split("\n");
    const { start, end } = node.loc;
    if (start.line === end.line) {
      return lines[start.line - 1].slice(start.column, end.column);
    }
    const parts = [lines[start.line - 1].slice(start.column)];
    for (let i = start.line; i < end.line - 1; i++) parts.push(lines[i]);
    parts.push(lines[end.line - 1].slice(0, end.column));
    return parts.join("\n");
  }

  function getLineIndent(code, lineNum) {
    const m = (code.split("\n")[lineNum - 1] || "").match(/^(\s*)/);
    return m ? m[1] : "  ";
  }

  function scoreByPatterns(code) {
    const scores = { array: 0, string: 0, tree: 0, graph: 0 };
    for (const re of TREE_SIGNALS) if (re.test(code)) scores.tree += 2;
    for (const re of GRAPH_SIGNALS) if (re.test(code)) scores.graph += 2;
    for (const re of STRING_SIGNALS) if (re.test(code)) scores.string += 2;
    for (const re of ARRAY_SIGNALS) if (re.test(code)) scores.array += 2;
    if (!scores.array && !scores.string && !scores.tree && !scores.graph) scores.array = 1;
    return scores;
  }

  function detectStructureType(code, ast) {
    const scores = scoreByPatterns(code);
    walk(ast, (node) => {
      if (node.type === "MemberExpression" && node.computed) {
        const obj = getSource(code, node.object);
        if (/\w+\[\s*\w+\s*\]\[\s*\w+\s*\]/.test(getSource(code, node)) || /grid|image|board|matrix/i.test(obj)) scores.graph += 4;
        else if (/^(s|str|chars|string)/i.test(obj.trim())) scores.string += 3;
        else scores.array += 1;
      }
      if (node.type === "MemberExpression" && !node.computed) {
        const prop = node.property.name || "";
        if (prop === "left" || prop === "right") scores.tree += 3;
      }
      if (node.type === "CallExpression") {
        const callee = getSource(code, node.callee);
        if (/buildTreeFromLevelOrder/.test(callee)) scores.tree += 6;
        if (/parseGraphInput/.test(callee)) scores.graph += 6;
        if (/parseStringInput/.test(callee)) scores.string += 6;
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
      if (node.type === "IfStatement") {
        const cond = getSource(code, node.test);
        if (/\w+\[\s*\w+\s*\]/.test(cond)) patterns.add("compare");
        if (/\w+\[\s*\w+\s*\]\[\s*\w+\s*\]/.test(cond)) patterns.add("grid-access");
      }
      if (node.type === "MemberExpression" && !node.computed) {
        const prop = node.property.name || "";
        if (prop === "left" || prop === "right") patterns.add("tree-traverse");
      }
    });
    for (const name of POINTER_NAMES) {
      if (new RegExp(`(?:let|const|var)\\s+${name}\\b`).test(code)) {
        patterns.add("pointers");
        if (name === "left" || name === "right") patterns.add("two-pointers");
        if (name === "mid") patterns.add("binary-search");
        if (name === "r" || name === "c") patterns.add("grid-coords");
      }
    }
    if (/viz\.step|await\s+viz\./.test(code)) patterns.add("manual-viz");
    if (/queue|bfs|dfs/i.test(code)) patterns.add("graph-traversal");
    if (/stack/i.test(code)) patterns.add("stack");
    return [...patterns];
  }

  function findPrimaryDataVar(code, ast, structureType) {
    const candidates = [];
    walk(ast, (node) => {
      if (node.type !== "VariableDeclarator" || !node.id?.name || !node.init) return;
      const name = node.id.name;
      const initSrc = getSource(code, node.init);
      if (/\[\s*\.\.\.\s*data\s*\]/.test(initSrc) || initSrc.trim() === "data") {
        candidates.push({ name, priority: 10, endLine: node.loc.end.line });
      } else if (/String\s*\(\s*data/.test(initSrc)) {
        candidates.push({ name, priority: 10, endLine: node.loc.end.line });
      } else if (/buildTreeFromLevelOrder/.test(initSrc)) {
        candidates.push({ name, priority: 10, endLine: node.loc.end.line });
      } else if (/data\.map|parseGraphInput/.test(initSrc)) {
        candidates.push({ name, priority: 9, endLine: node.loc.end.line });
      }
    });
    const defaults = { array: "arr", string: "s", tree: "root", graph: "grid" };
    if (!candidates.length) {
      for (const name of ["arr", "nums", "s", "str", "grid", "image", "root", "tree"]) {
        if (new RegExp(`(?:const|let|var)\\s+${name}\\b`).test(code)) {
          candidates.push({ name, priority: 5, endLine: 1 });
        }
      }
    }
    candidates.sort((a, b) => b.priority - a.priority);
    const pick = candidates[0] || { name: defaults[structureType] || "arr", endLine: 1 };
    return { name: pick.name, primaryVarLine: pick.endLine + 1 };
  }

  function findRunFunction(ast) {
    let runFn = null;
    walk(ast, (node) => {
      if (node.type === "FunctionDeclaration" && node.id?.name === "run") runFn = node;
    });
    return runFn;
  }

  function hasManualViz(code) {
    return /viz\.step\s*\(|await\s+viz\.(step|setArray|setString|setTree|setGraph|highlight)/.test(code);
  }

  function ensureRunWrapper(code, analysis) {
    if (analysis.hasRunFunction || /function\s+run\s*\(/.test(code)) return code;

    const preambles = {
      array: "const arr = [...data];",
      string: "const s = String(data);",
      tree: "const root = buildTreeFromLevelOrder(data);",
      graph: "const grid = Array.isArray(data[0]) ? data.map((row) => row.map(Number)) : data;",
    };
    const preamble = preambles[analysis.structureType] || preambles.array;
    return `async function run(viz, data) {\n  ${preamble}\n${code}\n}`;
  }

  function analyze(code) {
    try {
      const ast = parseAst(code);
      let { structureType, confidence, scores } = detectStructureType(code, ast);
      const primary = findPrimaryDataVar(code, ast, structureType);
      walk(ast, (node) => {
        if (node.type !== "VariableDeclarator" || !node.init) return;
        const initSrc = getSource(code, node.init);
        if (/String\s*\(\s*data/.test(initSrc)) {
          structureType = "string";
          confidence = Math.max(confidence, 0.85);
        } else if (/buildTreeFromLevelOrder/.test(initSrc)) {
          structureType = "tree";
          confidence = Math.max(confidence, 0.85);
        } else if (/data\.map/.test(initSrc) && /row|grid|image/i.test(getSource(code, node.id))) {
          structureType = "graph";
          confidence = Math.max(confidence, 0.85);
        } else if (/\[\s*\.\.\.\s*data\s*\]/.test(initSrc)) {
          structureType = "array";
          confidence = Math.max(confidence, 0.85);
        }
      });
      const patterns = detectPatterns(code, ast);
      return {
        ok: true,
        ast,
        structureType,
        confidence,
        scores,
        patterns,
        primaryVar: primary.name,
        primaryVarLine: primary.primaryVarLine,
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
        primaryVarLine: 2,
        hasManualViz: hasManualViz(code),
        hasRunFunction: /function\s+run\b/.test(code),
      };
    }
  }

  function traceSnippet(indent, line, event, opts = {}) {
    const parts = [`line: ${line}`, `event: '${event}'`];
    if (opts.varNames?.length) parts.push(`vars: { ${opts.varNames.join(", ")} }`);
    if (opts.dataVar) parts.push(`data: ${opts.dataVar}`);
    if (opts.node) parts.push(`node: ${opts.node}`);
    if (opts.indices?.length) parts.push(`indices: [${opts.indices.join(", ")}]`);
    if (opts.cellVars?.length) parts.push(`cellVars: [${opts.cellVars.map((v) => `'${v}'`).join(", ")}]`);
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

  function extractGridCoords(condSrc) {
    const m = condSrc.match(/(\w+)\[\s*(\w+)\s*\]\[\s*(\w+)\s*\]/);
    if (m) return { grid: m[1], row: m[2], col: m[3] };
    return null;
  }

  function extractTreeNodeRef(src) {
    const assignMatch = src.match(/(\w+)\s*=\s*(\w+)\.(left|right)\b/);
    if (assignMatch) return assignMatch[2];
    const propMatch = src.match(/(\w+)\.(left|right)\b/);
    if (propMatch) return propMatch[1];
    return null;
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
    const struct = analysis.structureType;
    const insertions = [];

    insertions.push({
      line: analysis.primaryVarLine,
      text: traceSnippet(getLineIndent(code, analysis.primaryVarLine), analysis.primaryVarLine, "start", {
        dataVar: primary,
        message: `Initialize ${struct} visualization`,
      }),
      priority: 0,
    });

    walk(runFn.body, (node, insideNestedFn) => {
      if (insideNestedFn) return;

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
        const line = node.loc.start.line;
        const indent = getLineIndent(code, line);

        const gridCoords = extractGridCoords(condSrc);
        if (gridCoords && struct === "graph") {
          insertions.push({
            line,
            text: traceSnippet(indent, line, "visit-cell", {
              varNames: [gridCoords.row, gridCoords.col],
              dataVar: primary,
              cellVars: [gridCoords.row, gridCoords.col],
              highlightType: "highlight",
              message: condSrc.slice(0, 50),
            }),
            priority: 2,
          });
        }

        const memberMatches = [...condSrc.matchAll(/(\w+)\[\s*(\w+)\s*\]/g)];
        if (memberMatches.length && struct !== "graph") {
          const obj = memberMatches[0][1];
          const indexVars = [...new Set(memberMatches.map((m) => m[2]).filter((v) => /^\w+$/.test(v)))];
          const hl = condSrc.includes("!==") || condSrc.includes("!=") ? "mismatch" : "compare";
          insertions.push({
            line,
            text: traceSnippet(indent, line, "compare", {
              varNames: indexVars,
              dataVar: primary,
              indices: indexVars,
              highlightType: hl,
              message: condSrc.slice(0, 50),
            }),
            priority: 2,
          });
        }

        if (struct === "tree") {
          const treeRef = extractTreeNodeRef(condSrc);
          if (treeRef) {
            insertions.push({
              line,
              text: traceSnippet(indent, line, "visit-node", {
                node: treeRef,
                dataVar: primary,
                highlightType: "visited",
                message: "Check tree branch",
              }),
              priority: 2,
            });
          }
        }
      }

      if (node.type === "ExpressionStatement") {
        const exprSrc = getSource(code, node.expression);

        if (node.expression?.type === "AssignmentExpression" && node.expression.left?.type === "ArrayExpression") {
          const indent = getLineIndent(code, node.loc.start.line);
          insertions.push({
            line: node.loc.end.line + 1,
            text: traceSnippet(indent, node.loc.end.line + 1, "swap", { dataVar: primary, highlightType: "swap", message: "Swap elements" }),
            priority: 3,
          });
        }

        if (struct === "tree" && /(\w+)\s*=\s*\w+\.(left|right)\b/.test(exprSrc)) {
          const line = node.loc.start.line;
          const indent = getLineIndent(code, line);
          const visitNode = extractTreeNodeRef(exprSrc);
          if (visitNode) {
            insertions.push({
              line,
              text: traceSnippet(indent, line, "visit-node", {
                node: visitNode,
                dataVar: primary,
                highlightType: "visited",
                message: "Tree traversal step",
              }),
              priority: 2,
            });
          }
        }

        if (struct === "tree" && /\.\s*val\b/.test(exprSrc)) {
          const line = node.loc.start.line;
          const indent = getLineIndent(code, line);
          const valMatch = exprSrc.match(/(\w+)\.val\b/);
          if (valMatch) {
            insertions.push({
              line,
              text: traceSnippet(indent, line, "visit-node", {
                node: valMatch[1],
                dataVar: primary,
                highlightType: "found",
                message: "Visit node",
              }),
              priority: 2,
            });
          }
        }

        if (struct === "graph" && /\w+\[\s*\w+\s*\]\[\s*\w+\s*\]\s*=/.test(exprSrc)) {
          const line = node.loc.end.line + 1;
          const indent = getLineIndent(code, line);
          insertions.push({
            line,
            text: traceSnippet(indent, line, "update-grid", { dataVar: primary, message: "Grid updated" }),
            priority: 3,
          });
        }
      }

      if (struct === "graph" && node.type === "CallExpression") {
        const callee = getSource(code, node.callee);
        if (/queue\.push|stack\.push/.test(callee)) {
          const line = node.loc.start.line;
          const indent = getLineIndent(code, line);
          const argsSrc = getSource(code, node);
          const coordMatch = argsSrc.match(/\[\s*(\w+)\s*,\s*(\w+)/);
          if (coordMatch) {
            insertions.push({
              line,
              text: traceSnippet(indent, line, "visit-cell", {
                varNames: [coordMatch[1], coordMatch[2]],
                dataVar: primary,
                cellVars: [coordMatch[1], coordMatch[2]],
                highlightType: "path",
                message: "Enqueue cell",
              }),
              priority: 2,
            });
          }
        }
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
    let working = code;
    let analysis = analyze(working);

    if (!analysis.hasRunFunction) {
      working = ensureRunWrapper(working, analysis);
      analysis = analyze(working);
    }

    let instrumented = { code: working, instrumented: false, reason: "skipped" };

    if (options.autoInstrument !== false && !analysis.hasManualViz) {
      if (analysis.hasRunFunction || /function\s+run\s*\(/.test(working)) {
        instrumented = instrument(working, analysis);
        if (instrumented.instrumented) working = instrumented.code;
      }
    }

    return { analysis, code: working, instrumented: instrumented.instrumented, instrumentReason: instrumented.reason };
  }

  return { analyze, instrument, prepare, ensureRunWrapper };
})();

function createAutoViz(viz, config) {
  const { structureType } = config;

  function syncData(data) {
    if (data === undefined || data === null) return;
    if (Array.isArray(data)) {
      if (data.length && Array.isArray(data[0])) viz.setGraph(data.map((r) => [...r]));
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
        if (typeof v === "number" && ["i", "j", "left", "right", "mid", "k"].includes(k)) out.push(v);
      }
    }
    return [...new Set(out)];
  }

  return {
    async trace({ line, event, vars = {}, data, node, indices = [], cellVars = [], highlightType = "compare", message = "" }) {
      syncData(data);

      const ptr = {};
      for (const [k, v] of Object.entries(vars)) {
        if (typeof v === "number") ptr[k] = v;
      }
      if (Object.keys(ptr).length) viz.setPointers(ptr);

      if (structureType === "tree" && node && node.id) {
        viz.highlightNodes([node.id]);
      }

      if (structureType === "graph" && cellVars.length >= 2) {
        const r = vars[cellVars[0]];
        const c = vars[cellVars[1]];
        if (typeof r === "number" && typeof c === "number") {
          viz.highlightCells([[r, c]], highlightType || "highlight");
        }
      }

      if (structureType === "string" && typeof vars.left === "number" && typeof vars.right === "number") {
        viz.highlightRange(vars.left, vars.right, highlightType || "compare");
      }

      const idx = resolveHighlightIndices(vars, indices);
      if (idx.length && (structureType === "array" || structureType === "string")) {
        viz.highlight(idx, highlightType || event);
      }

      if (structureType === "array" && data !== undefined && Array.isArray(data) && !Array.isArray(data[0])) {
        viz.setArray([...data]);
      }
      if (structureType === "string" && data !== undefined && typeof data === "string") {
        viz.setString(data);
      }
      if (structureType === "graph" && data !== undefined && Array.isArray(data) && Array.isArray(data[0])) {
        viz.setGraph(data.map((r) => [...r]));
      }

      await viz.step({
        line,
        message: message || `Auto: ${event}`,
        variables: node ? { ...vars, nodeVal: node.val } : vars,
      });
    },
  };
}

if (typeof module !== "undefined") {
  module.exports = { CodeParser, createAutoViz };
}
