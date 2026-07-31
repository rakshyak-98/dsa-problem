// REFLEX DRILL 05 — Trees & Stacks
//
// RUN: go run ./drills/05_trees_stacks_reflex
//
// AFTER PASSING: stacks/easy/valid_parentheses.js
package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TODO: REFLEX — inorder traversal (iterative with stack preferred)
func inorderTraversal(root *TreeNode) []int {
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

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}},
	}
	assert("inorderTraversal", reflect.DeepEqual(inorderTraversal(tree), []int{2, 1, 4, 3}))
	assert("inorderTraversal nil", reflect.DeepEqual(inorderTraversal(nil), []int{}))

	assert("maxDepth", maxDepth(tree) == 3)
	assert("maxDepth nil", maxDepth(nil) == 0)

	assert("isValidParentheses true", isValidParentheses("()[]{}") == true)
	assert("isValidParentheses false", isValidParentheses("(]") == false)
	assert("isValidParentheses empty", isValidParentheses("") == true)
	assert("isValidParentheses open only", isValidParentheses("(") == false)

	assert("dailyTemperatures", reflect.DeepEqual(
		dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}),
		[]int{1, 1, 4, 2, 1, 1, 0, 0},
	))
	assert("dailyTemperatures single", reflect.DeepEqual(dailyTemperatures([]int{50}), []int{0}))

	fmt.Println("\nAll trees/stacks reflex drills passed.")
	fmt.Println("Primary: stacks/easy/valid_parentheses.js")
}
