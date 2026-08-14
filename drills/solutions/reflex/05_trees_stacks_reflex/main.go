// SOLUTION — Reflex 05 Trees & Stacks (peek after honest attempt)
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

func inorderTraversal(root *TreeNode) []int {
	out := []int{}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		out = append(out, cur.Val)
		cur = cur.Right
	}
	return out
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	out := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		out = append(out, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return out
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	out := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		out = append(out, cur.Val)
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}

func levelOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	out := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		out = append(out, cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return out
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return 1 + l
	}
	return 1 + r
}

func isValidParentheses(s string) bool {
	stack := []byte{}
	match := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != match[ch] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}

func dailyTemperatures(temps []int) []int {
	stack := []int{}
	out := make([]int, len(temps))
	for i, t := range temps {
		for len(stack) > 0 && t > temps[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			out[j] = i - j
		}
		stack = append(stack, i)
	}
	return out
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
	assert("isValidParentheses invalid", !isValidParentheses("(]"))
	assert("isValidParentheses empty", isValidParentheses(""))
	assert("isValidParentheses open only", !isValidParentheses("("))
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
