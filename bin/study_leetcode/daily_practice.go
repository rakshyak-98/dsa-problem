// Daily LeetCode practice sets — 10 full problems per weekday topic.
// Problems are fetched from the LeetCode GraphQL API and saved to drills/leetcode/daily.json.
//
// RUN:              go run .
// Full catalog:     go run . -- --catalog
// Today's set only: go run . -- --set
// Force refresh:    go run . -- --refresh
package main

import "time"

var practiceSets = []practiceSet{
	{
		day: "Monday", topic: "Arrays", reflex: "01_arrays_reflex",
		topicTags: []string{"array"},
		seedSlugs: []string{
			"two-sum", "best-time-to-buy-and-sell-stock", "contains-duplicate", "maximum-subarray",
			"product-of-array-except-self", "merge-sorted-array", "rotate-array", "move-zeroes",
			"plus-one", "majority-element", "single-number", "intersection-of-two-arrays",
			"remove-element", "find-pivot-index", "running-sum-of-1d-array", "find-all-numbers-disappeared-in-an-array",
		},
		warmup: "Before coding: name the array shape (sorted? in-place?) and whether you need a prefix or two indices.",
		suggested: []string{
			"Start with Two Sum and Best Time to Buy — map + single pass patterns",
			"Save in-place problems (Rotate, Move Zeroes) for when you're warmed up",
		},
	},
	{
		day: "Tuesday", topic: "Hashing", reflex: "02_hashing_reflex",
		topicTags: []string{"hash-table"},
		seedSlugs: []string{
			"valid-anagram", "group-anagrams", "top-k-frequent-elements", "longest-consecutive-sequence",
			"two-sum", "contains-duplicate", "first-unique-character-in-a-string", "isomorphic-strings",
			"happy-number", "intersection-of-two-arrays-ii", "ransom-note", "design-hashmap",
			"4sum-ii", "subarray-sum-equals-k", "longest-harmonious-subsequence",
		},
		warmup: "Ask: do I need counts, seen-before, or a canonical key for grouping?",
		suggested: []string{
			"Group Anagrams and Top K Frequent are interview staples",
			"Longest Consecutive is a great 'sort vs hash' tradeoff question",
		},
	},
	{
		day: "Wednesday", topic: "Two Pointers", reflex: "03_two_pointers_reflex",
		topicTags: []string{"two-pointers", "sliding-window"},
		seedSlugs: []string{
			"valid-palindrome", "two-sum-ii-input-array-is-sorted", "3sum", "container-with-most-water",
			"trapping-rain-water", "remove-duplicates-from-sorted-array", "move-zeroes",
			"squares-of-a-sorted-array", "longest-substring-without-repeating-characters",
			"minimum-size-subarray-sum", "sort-colors", "backspace-string-compare",
			"reverse-string", "remove-nth-node-from-end-of-list", "partition-labels",
		},
		warmup: "Sorted input → L/R. In-place filter → read/write. Subarray length → sliding window.",
		suggested: []string{
			"Container and 3Sum reward clean pointer discipline",
			"Longest Substring Without Repeating is the classic variable window",
		},
	},
	{
		day: "Thursday", topic: "Binary Search", reflex: "04_binary_search_reflex",
		topicTags: []string{"binary-search"},
		seedSlugs: []string{
			"binary-search", "search-insert-position", "find-minimum-in-rotated-sorted-array",
			"search-in-rotated-sorted-array", "first-bad-version", "sqrtx", "valid-perfect-square",
			"find-smallest-letter-greater-than-target", "peak-index-in-a-mountain-array",
			"koko-eating-bananas", "guess-number-higher-or-lower", "single-element-in-a-sorted-array",
			"find-peak-element", "capacity-to-ship-packages-within-d-days", "split-array-largest-sum",
		},
		warmup: "Exact find → lo <= hi. First >= target → lower bound. Answer space → BS on answer.",
		suggested: []string{
			"Search Insert and First Bad Version are warm-up templates",
			"Koko Eating Bananas introduces binary search on the answer",
		},
	},
	{
		day: "Friday", topic: "Trees & Stacks", reflex: "05_trees_stacks_reflex",
		topicTags: []string{"tree", "stack", "binary-tree"},
		seedSlugs: []string{
			"valid-parentheses", "binary-tree-inorder-traversal", "maximum-depth-of-binary-tree",
			"invert-binary-tree", "same-tree", "symmetric-tree", "path-sum", "daily-temperatures",
			"min-stack", "binary-tree-level-order-traversal", "evaluate-reverse-polish-notation",
			"next-greater-element-i", "diameter-of-binary-tree", "subtree-of-another-tree",
			"implement-stack-using-queues",
		},
		warmup: "Tree traversal → stack/recursion. Matching brackets → stack. Next greater → monotonic stack.",
		suggested: []string{
			"Valid Parentheses and Inorder Traversal mirror today's reflex drill",
			"Daily Temperatures is the monotonic stack template",
		},
	},
	{
		day: "Saturday", topic: "Dynamic Programming", reflex: "06_dp_reflex",
		topicTags: []string{"dynamic-programming"},
		seedSlugs: []string{
			"climbing-stairs", "house-robber", "min-cost-climbing-stairs", "coin-change",
			"longest-increasing-subsequence", "word-break", "decode-ways", "unique-paths",
			"maximum-product-subarray", "partition-equal-subset-sum", "best-time-to-buy-and-sell-stock",
			"perfect-squares", "word-break-ii", "house-robber-ii", "longest-common-subsequence",
		},
		warmup: "Define dp[i] in English first. Linear DP → look back 1–2 states. Grid → 2D table.",
		suggested: []string{
			"Climbing Stairs and House Robber are the 1D DP foundations",
			"Coin Change and Word Break test state definition under pressure",
		},
	},
	{
		day: "Sunday", topic: "Graphs", reflex: "07_graphs_reflex",
		topicTags: []string{"graph", "depth-first-search", "breadth-first-search"},
		seedSlugs: []string{
			"number-of-islands", "flood-fill", "clone-graph", "course-schedule", "rotting-oranges",
			"word-ladder", "surrounded-regions", "pacific-atlantic-water-flow", "network-delay-time",
			"find-if-path-exists-in-graph", "all-paths-from-source-to-target", "graph-valid-tree",
			"number-of-provinces", "keys-and-rooms", "open-the-lock",
		},
		warmup: "Grid component → DFS/BFS + visited. Shortest unweighted path → BFS. Dependencies → topo sort.",
		suggested: []string{
			"Number of Islands and Flood Fill match today's reflex patterns",
			"Course Schedule is the canonical cycle / topo sort question",
		},
	},
}

func todayMeta() practiceSet {
	return practiceSets[todayWeekdayIndex()]
}

func todaySet() practiceSet {
	return todayMeta()
}

func lcURL(slug string) string {
	return "https://leetcode.com/problems/" + slug + "/"
}

func todayWeekdayIndex() int {
	wd := int(time.Now().Weekday())
	return (wd + 6) % 7
}
