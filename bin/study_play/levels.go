package main

import (
	"sort"
	"time"
)

// Understanding levels
//
// Pattern recognition is not the same skill as recall. A function you can type
// from memory is not one you can *spot* inside a disguised problem statement.
// Every drilled function therefore carries a level, and the daily refresh picks
// the question form that matches the level you have actually earned:
//
//	L1 recall    — restate the ask, then write the shape from memory
//	L2 pattern   — read a cue with no function name in it and name the move
//	L3 transfer  — apply the move to a variant that was never drilled
//
// The level is derived from `.drill_log.json`, not declared by hand, so it
// tracks what you can do today rather than what you did once in August.
type level int

const (
	levelRecall level = iota + 1
	levelPattern
	levelTransfer
)

func (l level) label() string {
	switch l {
	case levelRecall:
		return "L1 recall"
	case levelPattern:
		return "L2 pattern"
	case levelTransfer:
		return "L3 transfer"
	}
	return "L?"
}

// reviewGap is how long a function at this level may sit untouched before the
// refresh treats it as due again. Weaker levels come back sooner.
func (l level) reviewGap() int {
	switch l {
	case levelRecall:
		return 1
	case levelPattern:
		return 3
	case levelTransfer:
		return 7
	}
	return 3
}

// patternCue is one question in three forms.
//
//	ask  — the L1 form: what the function must return, in one sentence
//	cue  — the L2 form: what a problem statement looks like when this applies
//	move — the answer to the cue: what the hands write
//	twist— the L3 form: a variant that uses the same move on different shape
type patternCue struct {
	fn    string
	drill string
	tier  level
	ask   string
	cue   string
	move  string
	twist string
}

var cues = []patternCue{
	// 01 — arrays & prefix
	{"reverseInPlace", "01_arrays_reflex", levelRecall,
		"reverse the array in place and return it",
		"mirror / flip an array without extra space",
		"L/R pointers, swap, walk inward",
		"reverse only the words of a sentence, in place"},
	{"indexOfMax", "01_arrays_reflex", levelRecall,
		"index of the largest value (first index on ties)",
		"single best element while scanning once",
		"track best value + its index in one pass",
		"index of the second largest, still one pass"},
	{"arraySum", "01_arrays_reflex", levelRecall,
		"total of all elements",
		"aggregate the whole array",
		"accumulator loop",
		"sum only elements at even indices"},
	{"rotateRight", "01_arrays_reflex", levelPattern,
		"shift every element right by k, wrapping around",
		"cyclic shift by k, k may exceed n",
		"k %= n, then reverse all / reverse first k / reverse rest",
		"rotate left by k using the same three reversals"},
	{"runningSum", "01_arrays_reflex", levelPattern,
		"array where out[i] = sum of nums[0..i]",
		"repeated range-sum queries, or 'running total'",
		"prefix array; range(i,j) = pre[j] - pre[i-1]",
		"count subarrays whose sum equals k (prefix + map)"},

	// 02 — hashing
	{"twoSum", "02_hashing_reflex", levelPattern,
		"indices of the two values that add to target",
		"a pair that sums to target, array not sorted",
		"map value→index; look up target-x before inserting x",
		"three values summing to target (sort + fix one + two pointers)"},
	{"containsDuplicate", "02_hashing_reflex", levelRecall,
		"true if any value appears more than once",
		"'has it been seen before?'",
		"set; return true on first re-insert",
		"duplicate within distance k (sliding set)"},
	{"frequencyMap", "02_hashing_reflex", levelRecall,
		"map each value to its number of occurrences",
		"counting, majority, top-k, 'how many times'",
		"m[x]++ in one pass",
		"top k frequent elements (counts + bucket by count)"},
	{"firstUniqueChar", "02_hashing_reflex", levelPattern,
		"index of the first character that never repeats",
		"'first non-repeating' anything",
		"two passes: count, then rescan in order",
		"first repeating character instead of first unique"},
	{"groupAnagrams", "02_hashing_reflex", levelTransfer,
		"group strings that are rearrangements of each other",
		"'same multiset of letters' / equivalence classes",
		"canonical key (sorted string or count[26]) → map key→group",
		"group numbers by digit-sum, same canonical-key move"},

	// 03 — two pointers & window
	{"removeDuplicates", "03_two_pointers_reflex", levelRecall,
		"in-place unique prefix of a sorted array, return its length",
		"in-place filter, 'return the new length'",
		"read pointer scans, write pointer keeps",
		"allow each value at most twice"},
	{"moveZeroes", "03_two_pointers_reflex", levelRecall,
		"push every zero to the end, keep the rest in order",
		"partition in place while preserving order",
		"write pointer compacts keepers, then pad the tail",
		"move all even numbers to the front"},
	{"maxArea", "03_two_pointers_reflex", levelPattern,
		"largest area between two lines: width × min(height)",
		"maximize something over a pair of endpoints",
		"L/R pointers; move the side that limits the answer",
		"trapping rain water, same shrink-the-limiter idea"},
	{"isPalindrome", "03_two_pointers_reflex", levelRecall,
		"reads the same both ways, alphanumerics only, case-insensitive",
		"symmetry check on a sequence",
		"L/R pointers walking inward, skipping junk",
		"valid palindrome after deleting at most one character"},
	{"maxSumSubarrayK", "03_two_pointers_reflex", levelPattern,
		"best sum over any contiguous window of exactly k",
		"'subarray of size k' — fixed length stated in the problem",
		"build first window, then +nums[i] -nums[i-k]",
		"longest window with at most k distinct (variable window)"},

	// 04 — binary search
	{"binarySearch", "04_binary_search_reflex", levelRecall,
		"index of target in a sorted array, else -1",
		"sorted input + 'find' — or O(log n) demanded",
		"lo<=hi, mid = lo+(hi-lo)/2, discard a half",
		"search a rotated sorted array for an exact value"},
	{"searchInsert", "04_binary_search_reflex", levelPattern,
		"index of target, or where it would be inserted",
		"'first index ≥ target', insert position, count of smaller",
		"lower bound: lo<hi, keep hi=mid, answer is lo",
		"upper bound / count of values in a range"},
	{"findMinRotated", "04_binary_search_reflex", levelTransfer,
		"smallest value in a rotated sorted array",
		"sorted but rotated — the pivot is the answer",
		"compare mid with hi to decide which half is sorted",
		"find how many times the array was rotated"},
	{"isTargetPresent", "04_binary_search_reflex", levelRecall,
		"true if target exists in sorted data",
		"membership in sorted data without building a set",
		"binary search, return found rather than index",
		"first letter strictly greater than target (wrap around)"},

	// 05 — trees & stacks
	{"inorderTraversal", "05_trees_stacks_reflex", levelRecall,
		"node values left → node → right",
		"BST in sorted order, kth smallest",
		"stack: push all lefts, pop, visit, go right",
		"kth smallest in a BST — stop the inorder walk early"},
	{"preorderTraversal", "05_trees_stacks_reflex", levelRecall,
		"node values node → left → right",
		"serialize / copy a tree top-down",
		"stack: pop, visit, push right then left",
		"serialize and deserialize a binary tree"},
	{"postorderTraversal", "05_trees_stacks_reflex", levelPattern,
		"node values left → right → node",
		"children must be resolved before the parent",
		"two stacks, or preorder mirrored then reversed",
		"delete leaves / compute subtree sums bottom-up"},
	{"levelOrderTraversal", "05_trees_stacks_reflex", levelPattern,
		"values row by row, top to bottom",
		"'level', 'row', 'depth', 'closest first'",
		"BFS queue; capture len(queue) once per level",
		"right side view — last node of each level"},
	{"maxDepth", "05_trees_stacks_reflex", levelRecall,
		"longest root-to-leaf node count",
		"a tree answer built from both subtrees",
		"1 + max(left, right) recursion",
		"is the tree height-balanced"},
	{"isValidParentheses", "05_trees_stacks_reflex", levelPattern,
		"true if every closer matches the newest unmatched opener",
		"nesting, matching pairs, undo-most-recent",
		"stack of openers; mismatch or leftovers → false",
		"minimum insertions to balance the string"},
	{"dailyTemperatures", "05_trees_stacks_reflex", levelTransfer,
		"days to wait for a warmer value, 0 if none",
		"'next greater / next warmer / next smaller'",
		"monotonic decreasing stack of indices",
		"largest rectangle in a histogram"},

	// 06 — dp
	{"fib", "06_dp_reflex", levelRecall,
		"nth Fibonacci number",
		"answer depends on the previous two answers",
		"two rolling variables, no array",
		"tribonacci — same roll, three variables"},
	{"climbStairs", "06_dp_reflex", levelRecall,
		"number of ways to reach step n taking 1 or 2",
		"'how many ways' along a line",
		"ways[i] = ways[i-1] + ways[i-2]",
		"ways with steps of 1, 2 or 3"},
	{"minCostClimbingStairs", "06_dp_reflex", levelPattern,
		"cheapest way past the last step",
		"'minimum cost' along a line of choices",
		"dp[i] = cost[i] + min(dp[i-1], dp[i-2])",
		"min path sum down a triangle"},
	{"rob", "06_dp_reflex", levelPattern,
		"max total from non-adjacent elements",
		"'cannot take two in a row'",
		"take = skip_prev + v; skip = max(prev take, prev skip)",
		"houses in a circle — run it twice, drop one end each time"},

	// 07 — graphs
	{"numIslands", "07_graphs_reflex", levelPattern,
		"count of connected land regions",
		"'regions', 'groups', 'connected components' on a grid",
		"scan cells; each unvisited start = one DFS/BFS + mark",
		"largest island area — same sweep, return the max size"},
	{"floodFill", "07_graphs_reflex", levelRecall,
		"recolor every cell connected to the start",
		"paint bucket, spread from a seed",
		"DFS/BFS from start, guard against the no-op recolor",
		"surrounded regions — flood inward from the border"},
	{"shortestPathGrid", "07_graphs_reflex", levelTransfer,
		"fewest steps from corner to corner through open cells",
		"'shortest' / 'minimum steps' with unweighted moves",
		"BFS with a distance layer — never DFS for shortest",
		"rotting oranges — multi-source BFS from every rotten cell"},

	// 08 — heaps (bonus)
	{"kthLargest", "08_heap_reflex", levelPattern,
		"the kth largest value in the array",
		"'kth largest/smallest', 'top k'",
		"min-heap of size k, or quickselect",
		"kth largest in a stream — keep the heap alive"},
	{"lastStoneWeight", "08_heap_reflex", levelRecall,
		"weight left after repeatedly smashing the two heaviest",
		"repeatedly take the current extreme",
		"max-heap, pop two, push the difference",
		"minimum cost to connect ropes — pop the two smallest"},
	{"mergeKSorted", "08_heap_reflex", levelTransfer,
		"one sorted sequence from k sorted sequences",
		"merge many sorted inputs / k-way anything",
		"heap holding one frontier element per list",
		"smallest range covering an element from each list"},

	// 09 — backtracking (bonus)
	{"subsets", "09_backtrack_reflex", levelPattern,
		"every subset of the input",
		"'every possible group', each element is in or out",
		"recurse(i): choose, recurse(i+1), un-choose",
		"subsets of a multiset without duplicate results"},
	{"permute", "09_backtrack_reflex", levelTransfer,
		"every ordering of the input",
		"'all arrangements', order matters",
		"used[] flags + recurse, un-mark on the way out",
		"permutations with duplicates — sort, skip equal siblings"},
	{"combine", "09_backtrack_reflex", levelPattern,
		"all k-sized picks from 1..n",
		"'choose k of n' where order does not matter",
		"recurse with a start index so picks stay increasing",
		"combination sum — reuse the same number, prune on total"},

	// 10 — math add-on
	{"gcd", "10_math_reflex", levelRecall,
		"greatest common divisor of a and b",
		"'evenly divides', reduce a fraction, align cycles",
		"Euclid: gcd(b, a%b) until b is 0",
		"gcd of a whole array — fold gcd across it"},
	{"lcm", "10_math_reflex", levelRecall,
		"smallest number both a and b divide",
		"two repeating cycles meeting again",
		"a / gcd(a,b) * b — divide first to dodge overflow",
		"first day three schedules coincide"},
	{"modPow", "10_math_reflex", levelPattern,
		"a^b mod m without overflowing",
		"huge exponent, 'answer modulo 1e9+7'",
		"square-and-multiply on the bits of b",
		"nCr mod prime via Fermat inverse"},
	{"nCr", "10_math_reflex", levelPattern,
		"number of ways to choose r from n",
		"count paths in a grid, 'how many ways to choose'",
		"multiplicative loop with k = min(k, n-k)",
		"unique paths with obstacles — falls back to DP"},
	{"isPrime", "10_math_reflex", levelRecall,
		"true if n has no divisor besides 1 and itself",
		"primality or factor checks",
		"trial divide up to sqrt(n)",
		"count primes below n — sieve instead of per-number tests"},
	{"powOfTwo", "10_math_reflex", levelRecall,
		"true if n is an exact power of two",
		"halving, doubling, bit tricks",
		"n > 0 && n&(n-1) == 0",
		"count the set bits of n"},
}

var cueByFn = func() map[string]patternCue {
	m := make(map[string]patternCue, len(cues))
	for _, c := range cues {
		m[c.fn] = c
	}
	return m
}()

// owningFunction maps an assert case name back to the function it exercises,
// so "maxSumSubarrayK window slide" counts toward maxSumSubarrayK.
func owningFunction(caseName string) (string, bool) {
	best := ""
	for fn := range cueByFn {
		if len(fn) > len(best) && hasPrefixWord(caseName, fn) {
			best = fn
		}
	}
	return best, best != ""
}

func hasPrefixWord(s, prefix string) bool {
	if len(s) < len(prefix) || s[:len(prefix)] != prefix {
		return false
	}
	return len(s) == len(prefix) || s[len(prefix)] == ' '
}

// mastery is what the log says about one function today.
type mastery struct {
	cue      patternCue
	passes   int
	fails    int
	lastPass string
	earned   level // the level this function has earned
	daysIdle int   // days since the last pass; -1 when never passed
}

func (m mastery) attempted() bool { return m.passes+m.fails > 0 }

func (m mastery) due() bool {
	if !m.attempted() || m.daysIdle < 0 {
		return true
	}
	return m.daysIdle >= m.earned.reviewGap()
}

// earnedLevel grades a function from its log record. A function is never
// credited above its intrinsic tier: typing arraySum ten times does not make
// it a transfer question.
func earnedLevel(c patternCue, passes, fails int, passedRecently bool) level {
	if passes == 0 {
		return levelRecall
	}
	total := passes + fails
	failRate := float64(fails) / float64(total)
	got := levelRecall
	switch {
	case passes >= 5 && failRate < 0.15 && passedRecently:
		got = levelTransfer
	case passes >= 2 && failRate < 0.35:
		got = levelPattern
	}
	if got > c.tier {
		got = c.tier
	}
	return got
}

func daysSince(date string) int {
	if date == "" {
		return -1
	}
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return -1
	}
	d := int(time.Since(t).Hours() / 24)
	if d < 0 {
		return 0
	}
	return d
}

// masteryReport grades every catalog function against the drill log.
func masteryReport(root string) []mastery {
	log := loadLog(root)
	out := make([]mastery, 0, len(cues))
	for _, c := range cues {
		rec := log.Functions[c.fn]
		idle := daysSince(rec.LastPass)
		m := mastery{
			cue:      c,
			passes:   rec.Passes,
			fails:    rec.Fails,
			lastPass: rec.LastPass,
			daysIdle: idle,
			earned:   earnedLevel(c, rec.Passes, rec.Fails, idle >= 0 && idle <= 14),
		}
		out = append(out, m)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].cue.fn < out[j].cue.fn })
	return out
}
