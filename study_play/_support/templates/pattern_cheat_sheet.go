// PATTERN CHEAT SHEET — Templates for all reflex drills
//
// HOW TO USE:
// 1) Read one section once with comments.
// 2) Hide file, then rewrite from memory.
// 3) Repeat until each function is automatic.
//
// This file is reference only (not meant to be run).
package templates

import (
	"container/heap"
	"sort"
)

// =============================================================================
// DRILL 01 — ARRAYS
// =============================================================================

func ReverseInPlaceTemplate(arr []int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}

func IndexOfMaxTemplate(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	best := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[best] {
			best = i
		}
	}
	return best
}

func ArraySumTemplate(arr []int) int {
	sum := 0
	for _, x := range arr {
		sum += x
	}
	return sum
}

func RotateRightTemplate(arr []int, k int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	k %= n
	if k == 0 {
		return arr
	}
	reverseRange := func(a []int, l, r int) {
		for l < r {
			a[l], a[r] = a[r], a[l]
			l++
			r--
		}
	}
	reverseRange(arr, 0, n-1)
	reverseRange(arr, 0, k-1)
	reverseRange(arr, k, n-1)
	return arr
}

func RunningSumTemplate(arr []int) []int {
	out := make([]int, len(arr))
	sum := 0
	for i, x := range arr {
		sum += x
		out[i] = sum
	}
	return out
}

// =============================================================================
// DRILL 02 — HASHING
// =============================================================================

func TwoSumTemplate(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index
	for i, x := range nums {
		need := target - x
		if j, ok := seen[need]; ok {
			return []int{j, i}
		}
		seen[x] = i
	}
	return []int{}
}

func ContainsDuplicateTemplate(nums []int) bool {
	seen := make(map[int]bool)
	for _, x := range nums {
		if seen[x] {
			return true
		}
		seen[x] = true
	}
	return false
}

func FrequencyMapTemplate(arr []string) map[string]int {
	freq := make(map[string]int)
	for _, s := range arr {
		freq[s]++
	}
	return freq
}

func FirstUniqueCharTemplate(s string) string {
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if freq[s[i]] == 1 {
			return string(s[i])
		}
	}
	return ""
}

func GroupAnagramsTemplate(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, w := range strs {
		b := []byte(w)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		key := string(b)
		groups[key] = append(groups[key], w)
	}
	out := make([][]string, 0, len(groups))
	for _, g := range groups {
		out = append(out, g)
	}
	return out
}

// =============================================================================
// DRILL 03 — TWO POINTERS & FIXED WINDOW
// =============================================================================

func RemoveDuplicatesTemplate(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	write := 1
	for read := 1; read < len(nums); read++ {
		if nums[read] != nums[read-1] {
			nums[write] = nums[read]
			write++
		}
	}
	return write
}

func MoveZeroesTemplate(nums []int) {
	write := 0
	for read := 0; read < len(nums); read++ {
		if nums[read] != 0 {
			nums[write], nums[read] = nums[read], nums[write]
			write++
		}
	}
}

func MaxAreaTemplate(heights []int) int {
	left, right := 0, len(heights)-1
	best := 0
	for left < right {
		width := right - left
		h := heights[left]
		if heights[right] < h {
			h = heights[right]
		}
		area := width * h
		if area > best {
			best = area
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return best
}

func IsPalindromeTemplate(s string) bool {
	isAlphaNum := func(c byte) bool {
		return (c >= 'a' && c <= 'z') ||
			(c >= 'A' && c <= 'Z') ||
			(c >= '0' && c <= '9')
	}
	toLower := func(c byte) byte {
		if c >= 'A' && c <= 'Z' {
			return c + ('a' - 'A')
		}
		return c
	}
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isAlphaNum(s[left]) {
			left++
		}
		for left < right && !isAlphaNum(s[right]) {
			right--
		}
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func MaxSumSubarrayKTemplate(nums []int, k int) int {
	if len(nums) < k || k <= 0 {
		return 0
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	best := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > best {
			best = sum
		}
	}
	return best
}

// =============================================================================
// DRILL 04 — BINARY SEARCH
// =============================================================================

func BinarySearchTemplate(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

func SearchInsertTemplate(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func FindMinRotatedTemplate(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return nums[lo]
}

func IsTargetPresentTemplate(nums []int, target int) bool {
	return BinarySearchTemplate(nums, target) != -1
}

// =============================================================================
// DRILL 05 — TREES & STACKS
// =============================================================================

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversalTemplate(root *TreeNode) []int {
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

func MaxDepthTemplate(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MaxDepthTemplate(root.Left)
	right := MaxDepthTemplate(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func IsValidParenthesesTemplate(s string) bool {
	match := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != match[c] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}

func DailyTemperaturesTemplate(temps []int) []int {
	res := make([]int, len(temps))
	stack := []int{} // indices; temps decreasing
	for i := 0; i < len(temps); i++ {
		for len(stack) > 0 && temps[i] > temps[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[j] = i - j
		}
		stack = append(stack, i)
	}
	return res
}

// =============================================================================
// DRILL 06 — 1D DP
// =============================================================================

func FibTemplate(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func MinCostClimbingStairsTemplate(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		one := dp[i-1] + cost[i-1]
		two := dp[i-2] + cost[i-2]
		if one < two {
			dp[i] = one
		} else {
			dp[i] = two
		}
	}
	return dp[n]
}

func RobTemplate(nums []int) int {
	prev2, prev1 := 0, 0 // dp[i-2], dp[i-1]
	for _, x := range nums {
		take := prev2 + x
		skip := prev1
		cur := skip
		if take > skip {
			cur = take
		}
		prev2 = prev1
		prev1 = cur
	}
	return prev1
}

func ClimbStairsTemplate(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// =============================================================================
// DRILL 07 — GRAPHS (GRID BFS/DFS)
// =============================================================================

func NumIslandsTemplate(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}
	return count
}

func FloodFillTemplate(image [][]int, sr, sc, color int) [][]int {
	orig := image[sr][sc]
	if orig == color {
		return image
	}
	rows, cols := len(image), len(image[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || image[r][c] != orig {
			return
		}
		image[r][c] = color
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	dfs(sr, sc)
	return image
}

func ShortestPathGridTemplate(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	if grid[0][0] == 1 || grid[rows-1][cols-1] == 1 {
		return -1
	}
	type cell struct {
		r, c int
		d    int
	}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := []cell{{0, 0, 0}}
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	visited[0][0] = true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.r == rows-1 && cur.c == cols-1 {
			return cur.d
		}
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr < 0 || nc < 0 || nr >= rows || nc >= cols {
				continue
			}
			if grid[nr][nc] == 1 || visited[nr][nc] {
				continue
			}
			visited[nr][nc] = true
			queue = append(queue, cell{nr, nc, cur.d + 1})
		}
	}
	return -1
}

// =============================================================================
// VARIANTS — medium pattern recognition
// =============================================================================

func TwoSumSortedTemplate(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

func MaxSubarraySumTemplate(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	best, cur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if cur+nums[i] > nums[i] {
			cur = cur + nums[i]
		} else {
			cur = nums[i]
		}
		if cur > best {
			best = cur
		}
	}
	return best
}

func LengthOfLongestSubstringTemplate(s string) int {
	last := make(map[byte]int)
	left, best := 0, 0
	for right := 0; right < len(s); right++ {
		ch := s[right]
		if idx, ok := last[ch]; ok && idx >= left {
			left = idx + 1
		}
		last[ch] = right
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

func ProductExceptSelfTemplate(nums []int) []int {
	n := len(nums)
	out := make([]int, n)
	for i := range out {
		out[i] = 1
	}
	prefix := 1
	for i := 0; i < n; i++ {
		out[i] = prefix
		prefix *= nums[i]
	}
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		out[i] *= suffix
		suffix *= nums[i]
	}
	return out
}

// =============================================================================
// DRILL 08 — HEAPS
// =============================================================================

func KthLargestTemplate(nums []int, k int) int {
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

func LastStoneWeightTemplate(stones []int) int {
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

func MergeKSortedTemplate(lists [][]int) []int {
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

// =============================================================================
// DRILL 09 — BACKTRACKING
// =============================================================================

func SubsetsTemplate(nums []int) [][]int {
	var out [][]int
	var path []int
	var dfs func(int)
	dfs = func(start int) {
		cp := append([]int(nil), path...)
		out = append(out, cp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return out
}

func PermuteTemplate(nums []int) [][]int {
	var out [][]int
	path := make([]int, len(nums))
	used := make([]bool, len(nums))
	var dfs func(int)
	dfs = func(depth int) {
		if depth == len(nums) {
			cp := append([]int(nil), path...)
			out = append(out, cp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path[depth] = nums[i]
			dfs(depth + 1)
			used[i] = false
		}
	}
	dfs(0)
	return out
}

func CombineTemplate(n, k int) [][]int {
	var out [][]int
	path := []int{}
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			cp := append([]int(nil), path...)
			out = append(out, cp)
			return
		}
		need := k - len(path)
		for i := start; i <= n; i++ {
			if n-i+1 < need {
				break
			}
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return out
}
