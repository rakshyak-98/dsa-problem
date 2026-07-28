/**
 * DOM renderers for array, tree, and graph visualizations.
 */
const Renderers = {
  render(canvas, structureType, snapshot) {
    canvas.innerHTML = "";
    if (!snapshot) {
      canvas.innerHTML = '<p class="placeholder">No visualization data.</p>';
      return;
    }

    switch (structureType) {
      case "array":
        this.renderArray(canvas, snapshot);
        break;
      case "string":
        this.renderString(canvas, snapshot);
        break;
      case "tree":
        this.renderTree(canvas, snapshot);
        break;
      case "graph":
        this.renderGraph(canvas, snapshot);
        break;
      default:
        canvas.innerHTML = '<p class="placeholder">Unknown structure type.</p>';
    }
  },

  renderArray(canvas, snapshot) {
    const { array, highlights, pointers, arrayView = "bars", highlightRanges = [] } = snapshot;
    if (!array || array.length === 0) {
      canvas.innerHTML = '<p class="placeholder">Empty array.</p>';
      return;
    }

    const wrapper = document.createElement("div");
    wrapper.className = "array-viz";

    if (arrayView === "cells" || arrayView === "both") {
      wrapper.appendChild(this.buildIndexedCells(array, highlights, pointers, highlightRanges, "array"));
    }

    if (arrayView === "bars" || arrayView === "both") {
      const maxVal = Math.max(...array.map(Number).filter((n) => !Number.isNaN(n)), 1);
      const bars = document.createElement("div");
      bars.className = "array-bars";

      array.forEach((val, i) => {
        const cell = document.createElement("div");
        cell.className = "array-cell";

        const pointerLabels = Object.entries(pointers || {})
          .filter(([, idx]) => idx === i)
          .map(([name]) => name);

        if (pointerLabels.length) {
          const ptr = document.createElement("div");
          ptr.className = "array-pointer";
          ptr.textContent = pointerLabels.join(", ");
          cell.appendChild(ptr);
        }

        const bar = document.createElement("div");
        bar.className = "array-bar";
        const num = Number(val);
        const h = Number.isNaN(num) ? 40 : Math.max(30, (num / maxVal) * 160);
        bar.style.height = `${h}px`;
        bar.textContent = val;

        const hlType = highlights?.[i];
        if (hlType) bar.classList.add(hlType);

        const idx = document.createElement("div");
        idx.className = "array-index";
        idx.textContent = i;

        cell.appendChild(bar);
        cell.appendChild(idx);
        bars.appendChild(cell);
      });

      wrapper.appendChild(bars);
    }

    canvas.appendChild(wrapper);
  },

  renderString(canvas, snapshot) {
    const { string, highlights, pointers, highlightRanges = [] } = snapshot;
    if (!string || string.length === 0) {
      canvas.innerHTML = '<p class="placeholder">Empty string.</p>';
      return;
    }

    const chars = [...string];
    const wrapper = document.createElement("div");
    wrapper.className = "string-viz";
    wrapper.appendChild(this.buildIndexedCells(chars, highlights, pointers, highlightRanges, "string"));

    const full = document.createElement("div");
    full.className = "string-full";
    full.textContent = `"${string}"`;
    wrapper.appendChild(full);

    canvas.appendChild(wrapper);
  },

  buildIndexedCells(items, highlights, pointers, highlightRanges, kind) {
    const row = document.createElement("div");
    row.className = kind === "string" ? "string-cells" : "array-cells";

    const rangeMap = {};
    for (const { start, end, type } of highlightRanges || []) {
      for (let i = start; i <= end; i++) {
        rangeMap[i] = type;
      }
    }

    items.forEach((val, i) => {
      const cell = document.createElement("div");
      cell.className = "indexed-cell";

      const pointerLabels = Object.entries(pointers || {})
        .filter(([, idx]) => idx === i)
        .map(([name]) => name);

      if (pointerLabels.length) {
        const ptr = document.createElement("div");
        ptr.className = "cell-pointer";
        ptr.textContent = pointerLabels.join(", ");
        cell.appendChild(ptr);
      }

      const box = document.createElement("div");
      box.className = "cell-box";
      box.textContent = val;

      const hlType = highlights?.[i] || rangeMap[i];
      if (hlType) box.classList.add(hlType);

      const idx = document.createElement("div");
      idx.className = "cell-index";
      idx.textContent = i;

      cell.appendChild(box);
      cell.appendChild(idx);
      row.appendChild(cell);
    });

    return row;
  },

  renderTree(canvas, snapshot) {
    const { tree, highlightedNodes } = snapshot;
    if (!tree) {
      canvas.innerHTML = '<p class="placeholder">Empty tree.</p>';
      return;
    }

    const layout = layoutTree(tree);
    const highlightSet = new Set(highlightedNodes || []);

    const wrapper = document.createElement("div");
    wrapper.className = "tree-viz";

    const svg = document.createElementNS("http://www.w3.org/2000/svg", "svg");
    svg.setAttribute("width", layout.width);
    svg.setAttribute("height", layout.height);
    svg.setAttribute("viewBox", `0 0 ${layout.width} ${layout.height}`);

    for (const edge of layout.edges) {
      const line = document.createElementNS("http://www.w3.org/2000/svg", "line");
      line.setAttribute("x1", edge.x1);
      line.setAttribute("y1", edge.y1 + 16);
      line.setAttribute("x2", edge.x2);
      line.setAttribute("y2", edge.y2 - 16);
      line.setAttribute("class", "tree-edge");
      svg.appendChild(line);
    }

    for (const node of layout.nodes) {
      const g = document.createElementNS("http://www.w3.org/2000/svg", "g");
      g.setAttribute("class", "tree-node" + (highlightSet.has(node.id) ? " highlight" : ""));

      const circle = document.createElementNS("http://www.w3.org/2000/svg", "circle");
      circle.setAttribute("cx", node.cx);
      circle.setAttribute("cy", node.cy);
      circle.setAttribute("r", 22);

      const text = document.createElementNS("http://www.w3.org/2000/svg", "text");
      text.setAttribute("x", node.cx);
      text.setAttribute("y", node.cy);
      text.textContent = node.val;

      g.appendChild(circle);
      g.appendChild(text);
      svg.appendChild(g);
    }

    wrapper.appendChild(svg);
    canvas.appendChild(wrapper);
  },

  renderGraph(canvas, snapshot) {
    const { graph, highlightedCells, cellTypes } = snapshot;
    if (!graph || !graph.length) {
      canvas.innerHTML = '<p class="placeholder">Empty graph grid.</p>';
      return;
    }

    const highlightSet = new Set((highlightedCells || []).map(([r, c]) => `${r},${c}`));
    const types = cellTypes || {};

    const wrapper = document.createElement("div");
    wrapper.className = "graph-viz";

    const table = document.createElement("table");
    table.className = "graph-grid";

    graph.forEach((row, r) => {
      const tr = document.createElement("tr");
      row.forEach((cell, c) => {
        const td = document.createElement("td");
        td.className = "graph-cell";
        const key = `${r},${c}`;
        const type = types[key];

        if (type) {
          td.classList.add(type);
        } else if (highlightSet.has(key)) {
          td.classList.add("highlight");
        } else if (cell === 1 || cell === "1") {
          td.classList.add("wall");
        } else {
          td.classList.add("walkable");
        }

        td.textContent = cell;
        tr.appendChild(td);
      });
      table.appendChild(tr);
    });

    wrapper.appendChild(table);

    const coords = document.createElement("div");
    coords.className = "graph-coords";
    coords.textContent = "Rows ↓  Cols →";
    wrapper.appendChild(coords);

    canvas.appendChild(wrapper);
  },
};

if (typeof module !== "undefined") {
  module.exports = { Renderers };
}
