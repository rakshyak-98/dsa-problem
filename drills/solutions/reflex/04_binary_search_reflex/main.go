// SOLUTION — Reflex 04 Binary Search (peek after honest attempt)
package main

import "fmt"

func binarySearch(nums []int, target int) int {
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

func searchInsert(nums []int, target int) int {
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

func findMinRotated(nums []int) int {
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

func minEatingSpeed(piles []int, h int) int {
	// The search space is the answer itself (1..max pile), not an index. feasible
	// is monotonic: if speed s finishes in time, so does every speed above it.
	hours := func(speed int) int {
		total := 0
		for _, p := range piles {
			total += (p + speed - 1) / speed
		}
		return total
	}
	lo, hi := 1, 0
	for _, p := range piles {
		if p > hi {
			hi = p
		}
	}
	for lo < hi {
		mid := lo + (hi-lo)/2
		if hours(mid) <= h {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// quickSelect — quicksort that recurses only into the side holding rank k.
// Average O(n); the interview answer to "kth largest without a full sort".
func quickSelect(nums []int, k int) int {
	a := make([]int, len(nums))
	copy(a, nums)
	lo, hi, target := 0, len(a)-1, k-1
	for {
		p := partition(a, lo, hi)
		switch {
		case p == target:
			return a[p]
		case p < target:
			lo = p + 1
		default:
			hi = p - 1
		}
	}
}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}

// fastPow — binary exponentiation. Fold n bit by bit: square the base each
// step, multiply it into the result when the low bit is set.
func fastPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	result := 1.0
	for n > 0 {
		if n&1 == 1 {
			result *= x
		}
		x *= x
		n >>= 1
	}
	return result
}

// gcd — Euclid: gcd(a, b) = gcd(b, a mod b), stopping when b hits 0.
func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func almostEqual(a, b float64) bool {
	d := a - b
	if d < 0 {
		d = -d
	}
	return d < 1e-6
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
	}
	fmt.Printf("PASS: %s\n", name)
}

func main() {
	// binarySearch — hit/miss, empty, single, first/last/mid, two-element
	assert("binarySearch found mid", binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9) == 4)
	assert("binarySearch missing mid", binarySearch([]int{-1, 0, 3, 5, 9, 12}, 2) == -1)
	assert("binarySearch empty", binarySearch([]int{}, 1) == -1)
	assert("binarySearch single hit", binarySearch([]int{5}, 5) == 0)
	assert("binarySearch single miss", binarySearch([]int{5}, 2) == -1)
	assert("binarySearch first", binarySearch([]int{1, 2, 3, 4, 5}, 1) == 0)
	assert("binarySearch last", binarySearch([]int{1, 2, 3, 4, 5}, 5) == 4)
	assert("binarySearch two hit", binarySearch([]int{1, 3}, 3) == 1)
	assert("binarySearch two miss", binarySearch([]int{1, 3}, 2) == -1)

	// searchInsert — exist, middle, end, start, empty, single, equal front
	assert("searchInsert exist", searchInsert([]int{1, 3, 5, 6}, 5) == 2)
	assert("searchInsert new mid", searchInsert([]int{1, 3, 5, 6}, 2) == 1)
	assert("searchInsert end", searchInsert([]int{1, 3, 5, 6}, 7) == 4)
	assert("searchInsert empty", searchInsert([]int{}, 5) == 0)
	assert("searchInsert start", searchInsert([]int{2, 4, 6}, 1) == 0)
	assert("searchInsert single", searchInsert([]int{5}, 5) == 0)
	assert("searchInsert after single", searchInsert([]int{5}, 7) == 1)

	// findMinRotated — pivot, two elem, sorted, single, pivot at end
	assert("findMinRotated pivot", findMinRotated([]int{4, 5, 6, 7, 0, 1, 2}) == 0)
	assert("findMinRotated two", findMinRotated([]int{3, 1}) == 1)
	assert("findMinRotated sorted", findMinRotated([]int{1, 2, 3, 4}) == 1)
	assert("findMinRotated single", findMinRotated([]int{2}) == 2)
	assert("findMinRotated pivot end", findMinRotated([]int{2, 3, 4, 5, 1}) == 1)

	// minEatingSpeed — exact fit, one pile, h == len(piles), huge h
	assert("minEatingSpeed basic", minEatingSpeed([]int{3, 6, 7, 11}, 8) == 4)
	assert("minEatingSpeed tight", minEatingSpeed([]int{30, 11, 23, 4, 20}, 5) == 30)
	assert("minEatingSpeed loose", minEatingSpeed([]int{30, 11, 23, 4, 20}, 6) == 23)
	assert("minEatingSpeed single pile", minEatingSpeed([]int{12}, 3) == 4)
	assert("minEatingSpeed huge h", minEatingSpeed([]int{1, 1, 1}, 100) == 1)
	assert("minEatingSpeed h equals piles", minEatingSpeed([]int{5, 5, 5}, 3) == 5)

	// quickSelect — smallest, kth, median, largest, duplicates
	assert("quickSelect smallest", quickSelect([]int{3, 2, 1, 5, 6, 4}, 1) == 1)
	assert("quickSelect median", quickSelect([]int{7, 10, 4, 3, 20, 15}, 3) == 7)
	assert("quickSelect largest", quickSelect([]int{7, 10, 4, 3, 20, 15}, 6) == 20)
	assert("quickSelect duplicates", quickSelect([]int{2, 2, 2, 2}, 3) == 2)

	// fastPow — square, zero/one/negative exponent, fractional base
	assert("fastPow square", almostEqual(fastPow(2, 10), 1024))
	assert("fastPow zero exponent", almostEqual(fastPow(5, 0), 1))
	assert("fastPow negative exponent", almostEqual(fastPow(2, -2), 0.25))
	assert("fastPow fractional base", almostEqual(fastPow(2.1, 3), 9.261))

	// gcd — basic, coprime, divides, with zero
	assert("gcd basic", gcd(12, 18) == 6)
	assert("gcd coprime", gcd(7, 13) == 1)
	assert("gcd with zero", gcd(0, 5) == 5)

	fmt.Println("\nAll binary search reflex drills passed.")
	fmt.Println("Primary: binary_search/easy/search_insertion_position.js")
}
