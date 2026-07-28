// PATTERN CHEAT SHEET — Muscle memory templates
//
// HOW TO USE:
// 1. Read each template once with comments.
// 2. Cover the file. Re-type from memory into a blank file.
// 3. Repeat daily until you can write any section in under 2 minutes.
//
// Do NOT run this file for logic — it's reference + blank practice skeletons.
package templates

// =============================================================================
// 1. HASH MAP — Two Sum / complement lookup
// Time: O(n)  |  Space: O(n)
// Trigger: "pair that sums to target", "complement exists"
// =============================================================================
func TwoSumTemplate(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index
	for i, n := range nums {
		need := target - n
		if j, ok := seen[need]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return []int{}
}

// =============================================================================
// 2. FREQUENCY MAP — count occurrences
// Trigger: "most frequent", "anagram", "duplicate"
// =============================================================================
func BuildFreqMap(arr []int) map[int]int {
	freq := make(map[int]int)
	for _, x := range arr {
		freq[x]++
	}
	return freq
}

// =============================================================================
// 3. TWO POINTERS — opposite ends on sorted array
// Trigger: sorted + pair sum, palindrome, container area
// =============================================================================
func TwoPointersOpposite(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		// use nums[left], nums[right]
		// move left++ or right-- based on condition
		left++
		right--
	}
}

// =============================================================================
// 4. TWO POINTERS — same direction (fast/slow or write index)
// Trigger: remove duplicates in-place, move zeroes, partition
// =============================================================================
func TwoPointersSameDirection(nums []int) int {
	write := 0
	for read := 0; read < len(nums); read++ {
		if true /* keep nums[read] */ {
			nums[write] = nums[read]
			write++
		}
	}
	return write // new length
}

// =============================================================================
// 5. SLIDING WINDOW — fixed size k
// Trigger: "subarray of size k", "maximum average"
// =============================================================================
func SlidingWindowFixed(nums []int, k int) int {
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	best := windowSum
	for i := k; i < len(nums); i++ {
		windowSum += nums[i] - nums[i-k]
		if windowSum > best {
			best = windowSum
		}
	}
	return best
}

// =============================================================================
// 6. SLIDING WINDOW — variable size
// Trigger: "longest/shortest subarray where condition"
// =============================================================================
func SlidingWindowVariable(s string) int {
	left := 0
	best := 0
	// state := map[...]...  // window state
	for right := 0; right < len(s); right++ {
		// expand: include s[right] in state
		for false /* window invalid */ {
			// shrink: remove s[left] from state
			left++
		}
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

// =============================================================================
// 7. BINARY SEARCH — exact search
// Trigger: sorted array, find index, O(log n)
// =============================================================================
func BinarySearchExact(nums []int, target int) int {
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

// =============================================================================
// 8. BINARY SEARCH — lower bound (first >= target)
// =============================================================================
func LowerBound(nums []int, target int) int {
	lo, hi := 0, len(nums) // exclusive
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

// =============================================================================
// 9. PREFIX SUM
// Trigger: range sum query, subarray sum equals k
// =============================================================================
func BuildPrefix(nums []int) []int {
	prefix := []int{0}
	for _, x := range nums {
		prefix = append(prefix, prefix[len(prefix)-1]+x)
	}
	return prefix
	// sum(i..j) = prefix[j+1] - prefix[i]
}

// =============================================================================
// 10. BFS — grid shortest path / level order
// =============================================================================
func BfsGrid(grid [][]int, startR, startC int) {
	rows, cols := len(grid), len(grid[0])
	type cell struct{ r, c int }
	queue := []cell{{startR, startC}}
	visited := map[[2]int]bool{{startR, startC}: true}
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			key := [2]int{nr, nc}
			if nr < 0 || nc < 0 || nr >= rows || nc >= cols {
				continue
			}
			if visited[key] {
				continue
			}
			// if grid[nr][nc] blocks { continue }
			visited[key] = true
			queue = append(queue, cell{nr, nc})
		}
	}
}

// =============================================================================
// 11. DFS — grid / tree recursion
// =============================================================================
func DfsGrid(grid [][]int, r, c int, visited map[[2]int]bool) {
	rows, cols := len(grid), len(grid[0])
	if r < 0 || c < 0 || r >= rows || c >= cols {
		return
	}
	key := [2]int{r, c}
	if visited[key] {
		return
	}
	visited[key] = true
	DfsGrid(grid, r+1, c, visited)
	DfsGrid(grid, r-1, c, visited)
	DfsGrid(grid, r, c+1, visited)
	DfsGrid(grid, r, c-1, visited)
}

// =============================================================================
// 12. 1D DP — bottom-up
// Trigger: "minimum cost", "max ways", Fibonacci-style
// =============================================================================
func Dp1D(n int) int {
	dp := make([]int, n+1)
	// dp[0] = base
	// dp[1] = base
	for i := 2; i <= n; i++ {
		// dp[i] = recurrence using dp[i-1], dp[i-2], etc.
		_ = i
	}
	return dp[n]
}

// =============================================================================
// 13. MONOTONIC STACK — next greater element pattern
// =============================================================================
func NextGreaterTemplate(nums []int) []int {
	stack := []int{} // indices, decreasing values
	result := make([]int, len(nums))
	for i := range result {
		result[i] = -1
	}
	for i, n := range nums {
		for len(stack) > 0 && n > nums[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = n
		}
		stack = append(stack, i)
	}
	return result
}
