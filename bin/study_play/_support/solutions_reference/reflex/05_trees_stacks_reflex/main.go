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
	assert("inorderTraversal", reflect.DeepEqual(inorderTraversal(tree), []int{2, 1, 4, 3}))
	assert("inorderTraversal nil", reflect.DeepEqual(inorderTraversal(nil), []int{}))

	assert("maxDepth", maxDepth(tree) == 3)
	assert("maxDepth nil", maxDepth(nil) == 0)

	assert("isValidParentheses true", isValidParentheses("()[]{}"))
	assert("isValidParentheses false", !isValidParentheses("(]"))
	assert("isValidParentheses empty", isValidParentheses(""))
	assert("isValidParentheses open only", !isValidParentheses("("))

	assert("dailyTemperatures", reflect.DeepEqual(
		dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}),
		[]int{1, 1, 4, 2, 1, 1, 0, 0},
	))
	assert("dailyTemperatures single", reflect.DeepEqual(dailyTemperatures([]int{50}), []int{0}))

	fmt.Println("\nAll trees/stacks reflex drills passed.")
}
