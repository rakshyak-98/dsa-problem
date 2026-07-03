/**
 * REFLEX DRILL 05 — Trees & Stacks
 *
 * RUN: node study_play/drills/05_trees_stacks_reflex.js
 */

class TreeNode {
  constructor(val, left = null, right = null) {
    this.val = val;
    this.left = left;
    this.right = right;
  }
}

// TODO: REFLEX — inorder traversal (iterative with stack preferred)
function inorderTraversal(root) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — max depth of binary tree
function maxDepth(root) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — valid parentheses
function isValidParentheses(s) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — daily temperatures (next greater to the right) — monotonic stack
function dailyTemperatures(temps) {
  throw new Error("Implement from memory");
}

function assert(name, cond) {
  if (!cond) throw new Error(`FAIL: ${name}`);
  console.log(`PASS: ${name}`);
}

const tree = new TreeNode(1, new TreeNode(2), new TreeNode(3, new TreeNode(4), null));
assert("inorderTraversal", JSON.stringify(inorderTraversal(tree)) === "[2,1,4,3]");
assert("maxDepth", maxDepth(tree) === 3);
assert("isValidParentheses true", isValidParentheses("()[]{}") === true);
assert("isValidParentheses false", isValidParentheses("(]") === false);
assert(
  "dailyTemperatures",
  JSON.stringify(dailyTemperatures([73, 74, 75, 71, 69, 72, 76, 73])) ===
    "[1,1,4,2,1,1,0,0]"
);

console.log("\nAll trees/stacks reflex drills passed.");

module.exports = { TreeNode, inorderTraversal, maxDepth, isValidParentheses, dailyTemperatures };
