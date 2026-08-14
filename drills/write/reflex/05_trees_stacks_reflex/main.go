// REFLEX DRILL 05 — Trees & Stacks
//
// RUN: go run -C drills/write/reflex/05_trees_stacks_reflex .
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
	single := &TreeNode{Val: 5}
	leftSkew := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	// inorderTraversal — full, nil, single, left skew
	assert("inorderTraversal basic", reflect.DeepEqual(inorderTraversal(tree), []int{2, 1, 4, 3}))
	assert("inorderTraversal nil", reflect.DeepEqual(inorderTraversal(nil), []int{}))
	assert("inorderTraversal single", reflect.DeepEqual(inorderTraversal(single), []int{5}))
	assert("inorderTraversal left skew", reflect.DeepEqual(inorderTraversal(leftSkew), []int{3, 2, 1}))

	// preorderTraversal — full, nil, single, left skew
	assert("preorderTraversal basic", reflect.DeepEqual(preorderTraversal(tree), []int{1, 2, 3, 4}))
	assert("preorderTraversal nil", reflect.DeepEqual(preorderTraversal(nil), []int{}))
	assert("preorderTraversal single", reflect.DeepEqual(preorderTraversal(single), []int{5}))
	assert("preorderTraversal left skew", reflect.DeepEqual(preorderTraversal(leftSkew), []int{1, 2, 3}))

	// postorderTraversal — full, nil, single, left skew
	assert("postorderTraversal basic", reflect.DeepEqual(postorderTraversal(tree), []int{2, 4, 3, 1}))
	assert("postorderTraversal nil", reflect.DeepEqual(postorderTraversal(nil), []int{}))
	assert("postorderTraversal single", reflect.DeepEqual(postorderTraversal(single), []int{5}))
	assert("postorderTraversal left skew", reflect.DeepEqual(postorderTraversal(leftSkew), []int{3, 2, 1}))

	// levelOrderTraversal — full, nil, single, left skew
	assert("levelOrderTraversal basic", reflect.DeepEqual(levelOrderTraversal(tree), []int{1, 2, 3, 4}))
	assert("levelOrderTraversal nil", reflect.DeepEqual(levelOrderTraversal(nil), []int{}))
	assert("levelOrderTraversal single", reflect.DeepEqual(levelOrderTraversal(single), []int{5}))
	assert("levelOrderTraversal left skew", reflect.DeepEqual(levelOrderTraversal(leftSkew), []int{1, 2, 3}))

	// maxDepth — full, nil, single, left skew
	assert("maxDepth basic", maxDepth(tree) == 3)
	assert("maxDepth nil", maxDepth(nil) == 0)
	assert("maxDepth single", maxDepth(single) == 1)
	assert("maxDepth left skew", maxDepth(leftSkew) == 3)

	// isValidParentheses — valid types, invalid, empty, unmatched open/close, nested, interleaved
	assert("isValidParentheses basic", isValidParentheses("()[]{}") == true)
	assert("isValidParentheses invalid", isValidParentheses("(]") == false)
	assert("isValidParentheses empty", isValidParentheses("") == true)
	assert("isValidParentheses open only", isValidParentheses("(") == false)
	assert("isValidParentheses close only", isValidParentheses(")") == false)
	assert("isValidParentheses nested", isValidParentheses("((()))") == true)
	assert("isValidParentheses interleaved false", isValidParentheses("([)]") == false)
	assert("isValidParentheses mixed valid", isValidParentheses("{[()()]}") == true)

	// dailyTemperatures — classic, single, decreasing, equal, increasing, pair
	assert("dailyTemperatures basic", reflect.DeepEqual(
		dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}),
		[]int{1, 1, 4, 2, 1, 1, 0, 0},
	))
	assert("dailyTemperatures single", reflect.DeepEqual(dailyTemperatures([]int{50}), []int{0}))
	assert("dailyTemperatures decreasing", reflect.DeepEqual(dailyTemperatures([]int{5, 4, 3}), []int{0, 0, 0}))
	assert("dailyTemperatures equal", reflect.DeepEqual(dailyTemperatures([]int{70, 70, 70}), []int{0, 0, 0}))
	assert("dailyTemperatures increasing", reflect.DeepEqual(dailyTemperatures([]int{60, 61, 62}), []int{1, 1, 0}))
	assert("dailyTemperatures pair", reflect.DeepEqual(dailyTemperatures([]int{55, 56}), []int{1, 0}))

	fmt.Println("\nAll trees/stacks reflex drills passed.")
	fmt.Println("Primary: stacks/easy/valid_parentheses.js")
}
