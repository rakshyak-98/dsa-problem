// SOLUTIONS — Reference implementations for reflex drills
//
// USE ONLY AFTER honest attempt (15+ min per file).
// Cover this file and re-implement in the drill files until automatic.
//
// Do not import this during practice.
// This file is reference only — not meant to be run.
package solutions

import "sort"

// ---------- 01 arrays ----------

func ReverseInPlace(arr []int) []int {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return arr
}

func IndexOfMax(arr []int) int {
	best := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[best] {
			best = i
		}
	}
	return best
}

func ArraySum(arr []int) int {
	sum := 0
	for _, x := range arr {
		sum += x
	}
	return sum
}

func RotateRight(arr []int, k int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	k %= n
	out := make([]int, n)
	copy(out, arr[n-k:])
	copy(out[k:], arr[:n-k])
	return out
}

func RunningSum(arr []int) []int {
	out := make([]int, len(arr))
	sum := 0
	for i, x := range arr {
		sum += x
		out[i] = sum
	}
	return out
}

// ---------- 02 hashing ----------

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, n := range nums {
		need := target - n
		if j, ok := seen[need]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return []int{}
}

func ContainsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := set[n]; ok {
			return true
		}
		set[n] = struct{}{}
	}
	return false
}

func FrequencyMap(arr []string) map[string]int {
	m := make(map[string]int)
	for _, x := range arr {
		m[x]++
	}
	return m
}

func FirstUniqueChar(s string) string {
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}
	for _, ch := range s {
		if freq[ch] == 1 {
			return string(ch)
		}
	}
	return ""
}

func GroupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, w := range strs {
		b := []byte(w)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		key := string(b)
		m[key] = append(m[key], w)
	}
	out := make([][]string, 0, len(m))
	for _, g := range m {
		out = append(out, g)
	}
	return out
}

// ---------- 03 two pointers ----------

func RemoveDuplicates(nums []int) int {
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

func MoveZeroes(nums []int) {
	write := 0
	for read := 0; read < len(nums); read++ {
		if nums[read] != 0 {
			nums[write] = nums[read]
			write++
		}
	}
	for write < len(nums) {
		nums[write] = 0
		write++
	}
}

func MaxArea(heights []int) int {
	l, r := 0, len(heights)-1
	best := 0
	for l < r {
		h := heights[l]
		if heights[r] < h {
			h = heights[r]
		}
		area := h * (r - l)
		if area > best {
			best = area
		}
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return best
}

func IsPalindrome(s string) bool {
	clean := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c = c + ('a' - 'A')
		}
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			clean = append(clean, c)
		}
	}
	l, r := 0, len(clean)-1
	for l < r {
		if clean[l] != clean[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func MaxSumSubarrayK(nums []int, k int) int {
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

// ---------- 04 binary search ----------

func BinarySearch(nums []int, target int) int {
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

func SearchInsert(nums []int, target int) int {
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

func FindMinRotated(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func IsTargetPresent(nums []int, target int) bool {
	return BinarySearch(nums, target) != -1
}

// ---------- 05 trees/stacks ----------

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
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

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := MaxDepth(root.Left)
	r := MaxDepth(root.Right)
	if l > r {
		return 1 + l
	}
	return 1 + r
}

func IsValidParentheses(s string) bool {
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

func DailyTemperatures(temps []int) []int {
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

// ---------- 06 dp ----------

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func MinCostClimbingStairs(cost []int) int {
	a, b := 0, 0
	for i := 2; i <= len(cost); i++ {
		c := b + cost[i-1]
		if a+cost[i-2] < c {
			c = a + cost[i-2]
		}
		a, b = b, c
	}
	return b
}

func Rob(nums []int) int {
	prev2, prev1 := 0, 0
	for _, x := range nums {
		cur := prev1
		if prev2+x > cur {
			cur = prev2 + x
		}
		prev2, prev1 = prev1, cur
	}
	return prev1
}

func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// ---------- 07 graphs ----------

func NumIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	count := 0

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

func FloodFill(image [][]int, sr, sc, color int) [][]int {
	start := image[sr][sc]
	if start == color {
		return image
	}
	rows, cols := len(image), len(image[0])
	stack := [][2]int{{sr, sc}}
	for len(stack) > 0 {
		cell := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		r, c := cell[0], cell[1]
		if r < 0 || c < 0 || r >= rows || c >= cols || image[r][c] != start {
			continue
		}
		image[r][c] = color
		stack = append(stack, [2]int{r + 1, c}, [2]int{r - 1, c}, [2]int{r, c + 1}, [2]int{r, c - 1})
	}
	return image
}

func ShortestPathGrid(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	if grid[0][0] == 1 || grid[rows-1][cols-1] == 1 {
		return -1
	}
	type cell struct{ r, c, dist int }
	queue := []cell{{0, 0, 1}}
	grid[0][0] = 1
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.r == rows-1 && cur.c == cols-1 {
			return cur.dist
		}
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr < 0 || nc < 0 || nr >= rows || nc >= cols || grid[nr][nc] == 1 {
				continue
			}
			grid[nr][nc] = 1
			queue = append(queue, cell{nr, nc, cur.dist + 1})
		}
	}
	return -1
}
