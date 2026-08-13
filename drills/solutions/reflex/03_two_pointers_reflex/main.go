// SOLUTION — Reflex 03 Two Pointers (peek after honest attempt)
package main

import (
	"fmt"
	"reflect"
)

func removeDuplicates(nums []int) int {
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

func moveZeroes(nums []int) {
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

func maxArea(heights []int) int {
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

func isPalindrome(s string) bool {
	clean := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
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

func maxSumSubarrayK(nums []int, k int) int {
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

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	dup := []int{1, 1, 2, 2, 3}
	assert("removeDuplicates len", removeDuplicates(dup) == 3)
	assert("removeDuplicates vals", reflect.DeepEqual(dup[:3], []int{1, 2, 3}))
	assert("removeDuplicates empty", removeDuplicates([]int{}) == 0)
	singleDup := []int{7}
	assert("removeDuplicates single", removeDuplicates(singleDup) == 1 && singleDup[0] == 7)
	allSame := []int{2, 2, 2}
	assert("removeDuplicates all same", removeDuplicates(allSame) == 1 && allSame[0] == 2)

	zeros := []int{0, 1, 0, 3, 12}
	moveZeroes(zeros)
	assert("moveZeroes", reflect.DeepEqual(zeros, []int{1, 3, 12, 0, 0}))
	moveZeroesSingle := []int{0}
	moveZeroes(moveZeroesSingle)
	assert("moveZeroes single", reflect.DeepEqual(moveZeroesSingle, []int{0}))
	allZeros := []int{0, 0, 0}
	moveZeroes(allZeros)
	assert("moveZeroes all zeros", reflect.DeepEqual(allZeros, []int{0, 0, 0}))
	noZeros := []int{1, 2, 3}
	moveZeroes(noZeros)
	assert("moveZeroes no zeros", reflect.DeepEqual(noZeros, []int{1, 2, 3}))

	assert("maxArea", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
	assert("maxArea two", maxArea([]int{1, 1}) == 1)
	assert("maxArea wide shallow", maxArea([]int{1, 2, 1}) == 2)

	assert("isPalindrome", isPalindrome("A man, a plan, a canal: Panama"))
	assert("isPalindrome false", !isPalindrome("race a car"))
	assert("isPalindrome empty", isPalindrome(""))
	assert("isPalindrome single", isPalindrome("a"))
	assert("isPalindrome digits", isPalindrome("121"))

	assert("maxSumSubarrayK", maxSumSubarrayK([]int{2, 1, 5, 1, 3, 2}, 3) == 9)
	assert("maxSumSubarrayK k=1", maxSumSubarrayK([]int{4, 2, 9}, 1) == 9)
	assert("maxSumSubarrayK k=len", maxSumSubarrayK([]int{1, 2, 3}, 3) == 6)
	assert("maxSumSubarrayK negatives", maxSumSubarrayK([]int{-1, -2, -3}, 2) == -3)

	fmt.Println("\nAll two-pointer reflex drills passed.")
}
