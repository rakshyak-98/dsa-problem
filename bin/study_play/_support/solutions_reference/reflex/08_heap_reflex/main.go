// SOLUTION — Reflex 08 Heaps (peek after honest attempt)
package main

import (
	"container/heap"
	"fmt"
)

func kthLargest(nums []int, k int) int {
	h := &intMinHeap{}
	heap.Init(h)
	for _, n := range nums {
		heap.Push(h, n)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

type intMinHeap []int

func (h intMinHeap) Len() int            { return len(h) }
func (h intMinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h intMinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intMinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func lastStoneWeight(stones []int) int {
	h := &intMaxHeap{}
	heap.Init(h)
	for _, s := range stones {
		heap.Push(h, s)
	}
	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		if a > b {
			heap.Push(h, a-b)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}

type intMaxHeap []int

func (h intMaxHeap) Len() int            { return len(h) }
func (h intMaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h intMaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intMaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

type mergeItem struct {
	val int
	i   int
	j   int
}

func mergeKSorted(lists [][]int) []int {
	h := &mergeItemHeap{}
	heap.Init(h)
	for i, list := range lists {
		if len(list) > 0 {
			heap.Push(h, mergeItem{list[0], i, 0})
		}
	}
	out := []int{}
	for h.Len() > 0 {
		cur := heap.Pop(h).(mergeItem)
		out = append(out, cur.val)
		if cur.j+1 < len(lists[cur.i]) {
			heap.Push(h, mergeItem{lists[cur.i][cur.j+1], cur.i, cur.j + 1})
		}
	}
	return out
}

type mergeItemHeap []mergeItem

func (h mergeItemHeap) Len() int            { return len(h) }
func (h mergeItemHeap) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h mergeItemHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *mergeItemHeap) Push(x interface{}) { *h = append(*h, x.(mergeItem)) }
func (h *mergeItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func reflectDeepEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	assert("kthLargest", kthLargest([]int{3, 2, 1, 5, 6, 4}, 2) == 5)
	assert("kthLargest dup", kthLargest([]int{3, 3, 3, 3}, 2) == 3)

	assert("lastStoneWeight", lastStoneWeight([]int{2, 7, 4, 1, 8, 1}) == 1)
	assert("lastStoneWeight single", lastStoneWeight([]int{5}) == 5)

	got := mergeKSorted([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}})
	assert("mergeKSorted", len(got) == 8 && got[0] == 1 && got[len(got)-1] == 6)
	assert("mergeKSorted empty", len(mergeKSorted([][]int{})) == 0)
	assert("mergeKSorted one list", reflectDeepEqual(mergeKSorted([][]int{{1, 2}}), []int{1, 2}))

	fmt.Println("\nAll heap reflex drills passed.")
}
