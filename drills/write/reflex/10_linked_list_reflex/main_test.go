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

// build makes a list from values; toSlice reads one back. Cycles are built by
// hand in TestHasCycle so toSlice is never called on one.
func build(vals ...int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

func toSlice(head *ListNode) []int {
	out := []int{}
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func TestReverseList(t *testing.T) {
	assert(t, "reverseList basic", reflect.DeepEqual(toSlice(reverseList(build(1, 2, 3, 4))), []int{4, 3, 2, 1}))
	assert(t, "reverseList two", reflect.DeepEqual(toSlice(reverseList(build(1, 2))), []int{2, 1}))
	assert(t, "reverseList single", reflect.DeepEqual(toSlice(reverseList(build(7))), []int{7}))
	assert(t, "reverseList empty", reverseList(nil) == nil)
	assert(t, "reverseList negatives", reflect.DeepEqual(toSlice(reverseList(build(-1, 0, 1))), []int{1, 0, -1}))
}

func TestHasCycle(t *testing.T) {
	assert(t, "hasCycle empty", !hasCycle(nil))
	assert(t, "hasCycle single no loop", !hasCycle(build(1)))
	assert(t, "hasCycle straight", !hasCycle(build(1, 2, 3, 4)))

	selfLoop := build(1)
	selfLoop.Next = selfLoop
	assert(t, "hasCycle self loop", hasCycle(selfLoop))

	tailToHead := build(1, 2, 3)
	tailToHead.Next.Next.Next = tailToHead
	assert(t, "hasCycle tail to head", hasCycle(tailToHead))

	tailToMiddle := build(1, 2, 3, 4, 5)
	tailToMiddle.Next.Next.Next.Next.Next = tailToMiddle.Next.Next
	assert(t, "hasCycle tail to middle", hasCycle(tailToMiddle))
}

func TestMiddleNode(t *testing.T) {
	// Even length returns the second middle.
	assert(t, "middleNode even", middleNode(build(1, 2, 3, 4)).Val == 3)
	assert(t, "middleNode odd", middleNode(build(1, 2, 3, 4, 5)).Val == 3)
	assert(t, "middleNode single", middleNode(build(9)).Val == 9)
	assert(t, "middleNode two", middleNode(build(1, 2)).Val == 2)
	assert(t, "middleNode empty", middleNode(nil) == nil)
}

func TestMergeTwoLists(t *testing.T) {
	assert(t, "mergeTwoLists basic", reflect.DeepEqual(toSlice(mergeTwoLists(build(1, 3, 5), build(2, 4, 6))), []int{1, 2, 3, 4, 5, 6}))
	assert(t, "mergeTwoLists uneven", reflect.DeepEqual(toSlice(mergeTwoLists(build(1, 2, 9), build(3))), []int{1, 2, 3, 9}))
	assert(t, "mergeTwoLists left empty", reflect.DeepEqual(toSlice(mergeTwoLists(nil, build(1, 2))), []int{1, 2}))
	assert(t, "mergeTwoLists right empty", reflect.DeepEqual(toSlice(mergeTwoLists(build(1, 2), nil)), []int{1, 2}))
	assert(t, "mergeTwoLists both empty", mergeTwoLists(nil, nil) == nil)
	assert(t, "mergeTwoLists duplicates", reflect.DeepEqual(toSlice(mergeTwoLists(build(1, 1), build(1, 1))), []int{1, 1, 1, 1}))
}

func TestRemoveNthFromEnd(t *testing.T) {
	assert(t, "removeNthFromEnd middle", reflect.DeepEqual(toSlice(removeNthFromEnd(build(1, 2, 3, 4, 5), 2)), []int{1, 2, 3, 5}))
	assert(t, "removeNthFromEnd last", reflect.DeepEqual(toSlice(removeNthFromEnd(build(1, 2, 3), 1)), []int{1, 2}))
	assert(t, "removeNthFromEnd head", reflect.DeepEqual(toSlice(removeNthFromEnd(build(1, 2, 3), 3)), []int{2, 3}))
	assert(t, "removeNthFromEnd only node", removeNthFromEnd(build(1), 1) == nil)
	assert(t, "removeNthFromEnd n too large", reflect.DeepEqual(toSlice(removeNthFromEnd(build(1, 2), 5)), []int{1, 2}))
}

func TestAll(t *testing.T) {
	t.Run("reverseList", TestReverseList)
	t.Run("hasCycle", TestHasCycle)
	t.Run("middleNode", TestMiddleNode)
	t.Run("mergeTwoLists", TestMergeTwoLists)
	t.Run("removeNthFromEnd", TestRemoveNthFromEnd)
	fmt.Println("\nAll linked list reflex drills passed.")
}
