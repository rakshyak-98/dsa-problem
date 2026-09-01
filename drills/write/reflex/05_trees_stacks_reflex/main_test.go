package main

import (
	"fmt"
	"reflect"
	"testing"
)

func assert(t *testing.T, name string, cond bool) {
	t.Helper()
	if !cond {
		fmt.Printf("FAIL: %s\n", name)
		t.Fatalf("FAIL: %s", name)
	}
	fmt.Printf("PASS: %s\n", name)
}

func sampleTree() *TreeNode {
	return &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}},
	}
}

func singleNode() *TreeNode { return &TreeNode{Val: 5} }

func leftSkewTree() *TreeNode {
	return &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
}

func TestInorderTraversal(t *testing.T) {
	assert(t, "inorderTraversal basic", reflect.DeepEqual(inorderTraversal(sampleTree()), []int{2, 1, 4, 3}))
	assert(t, "inorderTraversal nil", reflect.DeepEqual(inorderTraversal(nil), []int{}))
	assert(t, "inorderTraversal single", reflect.DeepEqual(inorderTraversal(singleNode()), []int{5}))
	assert(t, "inorderTraversal left skew", reflect.DeepEqual(inorderTraversal(leftSkewTree()), []int{3, 2, 1}))
}

func TestPreorderTraversal(t *testing.T) {
	assert(t, "preorderTraversal basic", reflect.DeepEqual(preorderTraversal(sampleTree()), []int{1, 2, 3, 4}))
	assert(t, "preorderTraversal nil", reflect.DeepEqual(preorderTraversal(nil), []int{}))
	assert(t, "preorderTraversal single", reflect.DeepEqual(preorderTraversal(singleNode()), []int{5}))
	assert(t, "preorderTraversal left skew", reflect.DeepEqual(preorderTraversal(leftSkewTree()), []int{1, 2, 3}))
}

func TestPostorderTraversal(t *testing.T) {
	assert(t, "postorderTraversal basic", reflect.DeepEqual(postorderTraversal(sampleTree()), []int{2, 4, 3, 1}))
	assert(t, "postorderTraversal nil", reflect.DeepEqual(postorderTraversal(nil), []int{}))
	assert(t, "postorderTraversal single", reflect.DeepEqual(postorderTraversal(singleNode()), []int{5}))
	assert(t, "postorderTraversal left skew", reflect.DeepEqual(postorderTraversal(leftSkewTree()), []int{3, 2, 1}))
}

func TestLevelOrderTraversal(t *testing.T) {
	assert(t, "levelOrderTraversal basic", reflect.DeepEqual(levelOrderTraversal(sampleTree()), []int{1, 2, 3, 4}))
	assert(t, "levelOrderTraversal nil", reflect.DeepEqual(levelOrderTraversal(nil), []int{}))
	assert(t, "levelOrderTraversal single", reflect.DeepEqual(levelOrderTraversal(singleNode()), []int{5}))
	assert(t, "levelOrderTraversal left skew", reflect.DeepEqual(levelOrderTraversal(leftSkewTree()), []int{1, 2, 3}))
}

func TestMaxDepth(t *testing.T) {
	assert(t, "maxDepth basic", maxDepth(sampleTree()) == 3)
	assert(t, "maxDepth nil", maxDepth(nil) == 0)
	assert(t, "maxDepth single", maxDepth(singleNode()) == 1)
	assert(t, "maxDepth left skew", maxDepth(leftSkewTree()) == 3)
}

func TestIsValidParentheses(t *testing.T) {
	assert(t, "isValidParentheses basic", isValidParentheses("()[]{}") == true)
	assert(t, "isValidParentheses invalid", isValidParentheses("(]") == false)
	assert(t, "isValidParentheses empty", isValidParentheses("") == true)
	assert(t, "isValidParentheses open only", isValidParentheses("(") == false)
	assert(t, "isValidParentheses close only", isValidParentheses(")") == false)
	assert(t, "isValidParentheses nested", isValidParentheses("((()))") == true)
	assert(t, "isValidParentheses interleaved false", isValidParentheses("([)]") == false)
	assert(t, "isValidParentheses mixed valid", isValidParentheses("{[()()]}") == true)
}

func TestDailyTemperatures(t *testing.T) {
	assert(t, "dailyTemperatures basic", reflect.DeepEqual(
		dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}),
		[]int{1, 1, 4, 2, 1, 1, 0, 0},
	))
	assert(t, "dailyTemperatures single", reflect.DeepEqual(dailyTemperatures([]int{50}), []int{0}))
	assert(t, "dailyTemperatures decreasing", reflect.DeepEqual(dailyTemperatures([]int{5, 4, 3}), []int{0, 0, 0}))
	assert(t, "dailyTemperatures equal", reflect.DeepEqual(dailyTemperatures([]int{70, 70, 70}), []int{0, 0, 0}))
	assert(t, "dailyTemperatures increasing", reflect.DeepEqual(dailyTemperatures([]int{60, 61, 62}), []int{1, 1, 0}))
	assert(t, "dailyTemperatures pair", reflect.DeepEqual(dailyTemperatures([]int{55, 56}), []int{1, 0}))
}

func TestIsValidBST(t *testing.T) {
	assert(t, "isValidBST empty", isValidBST(nil))
	assert(t, "isValidBST single", isValidBST(&TreeNode{Val: 1}))
	valid := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	assert(t, "isValidBST valid", isValidBST(valid))
	swapped := &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}
	assert(t, "isValidBST swapped", !isValidBST(swapped))
	dupe := &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}
	assert(t, "isValidBST equal values", !isValidBST(dupe))
	deep := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	assert(t, "isValidBST deep violation", !isValidBST(deep))
}

func TestAll(t *testing.T) {
	t.Run("inorderTraversal", TestInorderTraversal)
	t.Run("preorderTraversal", TestPreorderTraversal)
	t.Run("postorderTraversal", TestPostorderTraversal)
	t.Run("levelOrderTraversal", TestLevelOrderTraversal)
	t.Run("maxDepth", TestMaxDepth)
	t.Run("isValidParentheses", TestIsValidParentheses)
	t.Run("dailyTemperatures", TestDailyTemperatures)
	t.Run("isValidBST", TestIsValidBST)
	fmt.Println("\nAll trees/stacks reflex drills passed.")
}
