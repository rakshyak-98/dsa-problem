/**
 * Step recording engine for DSA visualizer.
 * User code calls viz.step() to snapshot state; playback navigates snapshots.
 */
class VizEngine {
  constructor(structureType) {
    this.structureType = structureType;
    this.steps = [];
    this.currentStep = -1;
    this.isPlaying = false;
    this.playTimer = null;
    this.onStepChange = null;
    this.onComplete = null;

    this.state = {
      array: [],
      string: "",
      arrayView: "bars",
      highlightRanges: [],
      graph: [],
      tree: null,
      treeLayout: null,
      highlights: {},
      pointers: {},
      highlightedNodes: [],
      highlightedCells: [],
      cellTypes: {},
      variables: {},
      message: "",
      line: null,
    };
  }

  reset() {
    this.steps = [];
    this.currentStep = -1;
    this.stopPlayback();
    this.state = {
      array: [],
      string: "",
      arrayView: "bars",
      highlightRanges: [],
      graph: [],
      tree: null,
      treeLayout: null,
      highlights: {},
      pointers: {},
      highlightedNodes: [],
      highlightedCells: [],
      cellTypes: {},
      variables: {},
      message: "",
      line: null,
    };
  }

  cloneState() {
    return {
      array: [...this.state.array],
      string: this.state.string,
      arrayView: this.state.arrayView,
      highlightRanges: this.state.highlightRanges.map((r) => ({ ...r })),
      graph: this.state.graph.map((row) => [...row]),
      tree: this.state.tree ? JSON.parse(JSON.stringify(this.state.tree)) : null,
      treeLayout: this.state.treeLayout
        ? JSON.parse(JSON.stringify(this.state.treeLayout))
        : null,
      highlights: { ...this.state.highlights },
      pointers: { ...this.state.pointers },
      highlightedNodes: [...this.state.highlightedNodes],
      highlightedCells: this.state.highlightedCells.map(([r, c]) => [r, c]),
      cellTypes: { ...this.state.cellTypes },
      variables: { ...this.state.variables },
      message: this.state.message,
      line: this.state.line,
    };
  }

  recordStep(overrides = {}) {
    const snapshot = { ...this.cloneState(), ...overrides };
    this.steps.push(snapshot);
    return this.steps.length - 1;
  }

  setArray(arr) {
    this.state.array = [...arr];
  }

  setString(str) {
    this.state.string = String(str);
  }

  setArrayView(view) {
    if (["bars", "cells", "both"].includes(view)) {
      this.state.arrayView = view;
    }
  }

  highlightRange(start, end, type = "compare") {
    this.state.highlightRanges = [{ start, end, type }];
  }

  addHighlightRange(start, end, type = "compare") {
    this.state.highlightRanges.push({ start, end, type });
  }

  clearHighlightRanges() {
    this.state.highlightRanges = [];
  }

  setGraph(grid) {
    this.state.graph = grid.map((row) => [...row]);
  }

  setTree(root) {
    this.state.tree = root;
  }

  setTreeLayout(layout) {
    this.state.treeLayout = layout;
  }

  highlight(indices, type = "compare") {
    const map = {};
    for (const i of indices) {
      map[i] = type;
    }
    this.state.highlights = map;
  }

  setPointers(pointers) {
    this.state.pointers = { ...pointers };
  }

  highlightNodes(nodeIds) {
    this.state.highlightedNodes = [...nodeIds];
  }

  highlightCells(cells, type = "highlight") {
    this.state.highlightedCells = cells.map(([r, c]) => [r, c]);
    const cellTypes = {};
    for (const [r, c] of cells) {
      cellTypes[`${r},${c}`] = type;
    }
    this.state.cellTypes = cellTypes;
  }

  setCellTypes(cellTypes) {
    this.state.cellTypes = { ...cellTypes };
  }

  setVariables(vars) {
    this.state.variables = { ...this.state.variables, ...vars };
  }

  createVizAPI() {
    const self = this;

    return {
      setArray(arr) {
        self.setArray(arr);
      },
      setString(str) {
        self.setString(str);
      },
      setArrayView(view) {
        self.setArrayView(view);
      },
      highlightRange(start, end, type) {
        self.highlightRange(start, end, type);
      },
      addHighlightRange(start, end, type) {
        self.addHighlightRange(start, end, type);
      },
      clearHighlightRanges() {
        self.clearHighlightRanges();
      },
      setGraph(grid) {
        self.setGraph(grid);
      },
      setTree(root) {
        self.setTree(root);
      },
      highlight(indices, type) {
        self.highlight(indices, type);
      },
      setPointers(pointers) {
        self.setPointers(pointers);
      },
      highlightNodes(nodeIds) {
        self.highlightNodes(nodeIds);
      },
      highlightCells(cells, type) {
        self.highlightCells(cells, type);
      },
      setCellTypes(cellTypes) {
        self.setCellTypes(cellTypes);
      },
      setVariables(vars) {
        self.setVariables(vars);
      },
      async step(opts = {}) {
        const {
          message = "",
          line = null,
          highlights,
          pointers,
          variables,
          highlightedNodes,
          highlightedCells,
          cellTypes,
          array,
          string,
          arrayView,
          highlightRanges,
          graph,
          delay = 0,
        } = opts;

        if (array !== undefined) self.setArray(array);
        if (string !== undefined) self.setString(string);
        if (arrayView !== undefined) self.setArrayView(arrayView);
        if (highlightRanges !== undefined) {
          self.state.highlightRanges = highlightRanges.map((r) => ({ ...r }));
        }
        if (graph !== undefined) self.setGraph(graph);
        if (highlights !== undefined) {
          if (Array.isArray(highlights)) {
            self.highlight(highlights, opts.highlightType || "compare");
          } else {
            self.state.highlights = { ...highlights };
          }
        }
        if (pointers !== undefined) self.setPointers(pointers);
        if (variables !== undefined) self.setVariables(variables);
        if (highlightedNodes !== undefined) self.highlightNodes(highlightedNodes);
        if (highlightedCells !== undefined) {
          self.highlightCells(highlightedCells, opts.cellType || "highlight");
        }
        if (cellTypes !== undefined) self.setCellTypes(cellTypes);

        self.state.message = message;
        self.state.line = line;
        self.recordStep();

        if (delay > 0) {
          await new Promise((r) => setTimeout(r, delay));
        }
      },
    };
  }

  getCurrentSnapshot() {
    if (this.currentStep < 0 || this.currentStep >= this.steps.length) {
      return this.cloneState();
    }
    return this.steps[this.currentStep];
  }

  goToStep(index) {
    if (index < 0 || index >= this.steps.length) return false;
    this.currentStep = index;
    if (this.onStepChange) {
      this.onStepChange(this.getCurrentSnapshot(), index, this.steps.length);
    }
    return true;
  }

  nextStep() {
    return this.goToStep(this.currentStep + 1);
  }

  prevStep() {
    return this.goToStep(this.currentStep - 1);
  }

  play(intervalMs, onEnd) {
    if (this.isPlaying) return;
    if (this.steps.length <= 1) return false;

    // Restart from beginning when already at the last step
    if (this.currentStep >= this.steps.length - 1) {
      this.goToStep(0);
    }

    this.isPlaying = true;

    const tick = () => {
      if (!this.isPlaying) return;
      const atLast = this.currentStep >= this.steps.length - 1;
      if (atLast) {
        this.stopPlayback();
        if (onEnd) onEnd();
        return;
      }
      this.nextStep();
      this.playTimer = setTimeout(tick, intervalMs);
    };

    if (this.currentStep < 0 && this.steps.length > 0) {
      this.goToStep(0);
    }
    this.playTimer = setTimeout(tick, intervalMs);
  }

  stopPlayback() {
    this.isPlaying = false;
    if (this.playTimer) {
      clearTimeout(this.playTimer);
      this.playTimer = null;
    }
  }

  togglePlayback(intervalMs, onEnd) {
    if (this.isPlaying) {
      this.stopPlayback();
      return false;
    }
    if (this.steps.length <= 1) return false;
    this.play(intervalMs, onEnd);
    return true;
  }
}

/** Build binary tree from level-order array (null = empty). */
function buildTreeFromLevelOrder(values) {
  if (!values.length || values[0] == null) return null;

  const nodes = values.map((v, i) =>
    v == null || v === "null" ? null : { id: `n${i}`, val: Number(v), left: null, right: null }
  );

  const root = nodes[0];
  const queue = [root];
  let i = 1;

  while (queue.length && i < nodes.length) {
    const node = queue.shift();
    if (!node) continue;

    if (i < nodes.length && nodes[i]) {
      node.left = nodes[i];
      queue.push(nodes[i]);
    }
    i++;

    if (i < nodes.length && nodes[i]) {
      node.right = nodes[i];
      queue.push(nodes[i]);
    }
    i++;
  }

  return root;
}

/** Compute SVG layout positions for a binary tree. */
function layoutTree(root) {
  if (!root) return { nodes: [], edges: [], width: 200, height: 100 };

  const nodes = [];
  const edges = [];

  function walk(node, depth, pos, spread) {
    if (!node) return;
    nodes.push({ id: node.id, val: node.val, x: pos, y: depth });
    if (node.left) {
      edges.push({ from: node.id, to: node.left.id, x1: pos, y1: depth, x2: pos - spread, y2: depth + 1 });
      walk(node.left, depth + 1, pos - spread, spread / 2);
    }
    if (node.right) {
      edges.push({ from: node.id, to: node.right.id, x1: pos, y1: depth, x2: pos + spread, y2: depth + 1 });
      walk(node.right, depth + 1, pos + spread, spread / 2);
    }
  }

  walk(root, 0, 0, 8);

  const xs = nodes.map((n) => n.x);
  const minX = Math.min(...xs);
  const maxX = Math.max(...xs);
  const maxY = Math.max(...nodes.map((n) => n.y));

  const padX = 60;
  const padY = 50;
  const scaleX = 40;
  const scaleY = 70;

  const normNodes = nodes.map((n) => ({
    ...n,
    cx: (n.x - minX) * scaleX + padX,
    cy: n.y * scaleY + padY,
  }));

  const nodeMap = Object.fromEntries(normNodes.map((n) => [n.id, n]));

  const normEdges = edges.map((e) => ({
    x1: nodeMap[e.from].cx,
    y1: nodeMap[e.from].cy,
    x2: nodeMap[e.to].cx,
    y2: nodeMap[e.to].cy,
  }));

  const width = (maxX - minX) * scaleX + padX * 2 + 40;
  const height = maxY * scaleY + padY * 2 + 40;

  return { nodes: normNodes, edges: normEdges, width, height };
}

if (typeof module !== "undefined") {
  module.exports = { VizEngine, buildTreeFromLevelOrder, layoutTree };
}
