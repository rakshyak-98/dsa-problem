// REFLEX DRILL 05 — Trees & Stacks
//
// RUN: go run -C drills/write/reflex/05_trees_stacks_reflex .
//
// AFTER PASSING: stacks/easy/valid_parentheses.js
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TODO: REFLEX — inorder traversal (iterative with stack preferred)
func inorderTraversal(root *TreeNode) []int {
	panic("Implement from memory")
}

// TODO: REFLEX — preorder traversal (iterative with stack preferred)
func preorderTraversal(root *TreeNode) []int {
	panic("Implement from memory")
}

// TODO: REFLEX — postorder traversal (iterative with stack preferred)
func postorderTraversal(root *TreeNode) []int {
	panic("Implement from memory")
}

// TODO: REFLEX — level-order traversal (BFS with queue)
func levelOrderTraversal(root *TreeNode) []int {
	panic("Implement from memory")
}

// TODO: REFLEX — max depth of binary tree
func maxDepth(root *TreeNode) int {
	panic("Implement from memory")
}

// TODO: REFLEX — valid parentheses
func isValidParentheses(s string) bool {
	panic("Implement from memory")
}

// TODO: REFLEX — daily temperatures (next greater to the right) — monotonic stack
func dailyTemperatures(temps []int) []int {
	panic("Implement from memory")
}

func main() {}
