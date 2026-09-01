// REFLEX DRILL 10 — Linked Lists
//
// RUN: go run -C drills/write/reflex/10_linked_list_reflex .
//
// AFTER PASSING: https://leetcode.com/problems/reverse-linked-list/
package main

// ListNode is given — the drill is the traversal, not the type.
type ListNode struct {
	Val  int
	Next *ListNode
}

// TODO: REFLEX — reverse the list, return the new head
func reverseList(head *ListNode) *ListNode {
	panic("Implement from memory")
}

// TODO: REFLEX — does the list loop back on itself (Floyd, O(1) space)
func hasCycle(head *ListNode) bool {
	panic("Implement from memory")
}

// TODO: REFLEX — middle node; on even length return the second middle
func middleNode(head *ListNode) *ListNode {
	panic("Implement from memory")
}

// TODO: REFLEX — merge two sorted lists into one sorted list
func mergeTwoLists(a, b *ListNode) *ListNode {
	panic("Implement from memory")
}

// TODO: REFLEX — drop the nth node from the end, return the head
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	panic("Implement from memory")
}

func main() {}
