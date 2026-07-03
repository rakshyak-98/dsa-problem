/**
 * REFLEX DRILL 07 — Graphs (Grid BFS/DFS)
 *
 * RUN: node study_play/drills/07_graphs_reflex.js
 */

// TODO: REFLEX — count islands ('1' land, '0' water)
function numIslands(grid) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — flood fill from (sr, sc) with color
function floodFill(image, sr, sc, color) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — BFS shortest path length in unweighted grid (0=walkable, 1=wall)
// Return -1 if no path from top-left to bottom-right
function shortestPathGrid(grid) {
  throw new Error("Implement from memory");
}

function assert(name, cond) {
  if (!cond) throw new Error(`FAIL: ${name}`);
  console.log(`PASS: ${name}`);
}

const grid1 = [
  ["1", "1", "0"],
  ["0", "1", "0"],
  ["1", "0", "1"],
];
assert("numIslands", numIslands(grid1) === 3);

const img = [
  [1, 1, 1],
  [1, 1, 0],
  [1, 0, 1],
];
assert(
  "floodFill",
  JSON.stringify(floodFill(JSON.parse(JSON.stringify(img)), 1, 1, 2)) ===
    JSON.stringify([
      [2, 2, 2],
      [2, 2, 0],
      [2, 0, 1],
    ])
);

const pathGrid = [
  [0, 0, 0],
  [1, 1, 0],
  [0, 0, 0],
];
assert("shortestPathGrid", shortestPathGrid(pathGrid) === 4);

console.log("\nAll graph reflex drills passed.");

module.exports = { numIslands, floodFill, shortestPathGrid };
