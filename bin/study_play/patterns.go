package main

// Shared pattern catalog, review schedule and LeetCode cards.
//
// This file is copied verbatim into bin/study_play and bin/study_leetcode (they
// are separate Go modules with no go.work, like argutil.go). Edit one, then
// re-copy it to the other.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// pattern is one problem-solving pattern: the thing to recognise before coding.
// signal is the wording in a problem statement that points to it.
type pattern struct {
	id     string
	name   string
	signal string
}

var patterns = []pattern{
	{"hash-lookup", "hash map / set lookup", "\"seen before?\", a complement, counts, grouping by a key"},
	{"two-pointers", "two pointers from both ends", "sorted input, a pair or a palindrome; move the side that cannot help"},
	{"read-write", "read/write pointers (in-place filter)", "\"in place\", remove or move elements, return the new length"},
	{"fixed-window", "fixed-size sliding window", "a contiguous run of exactly k"},
	{"variable-window", "variable sliding window", "the longest/shortest contiguous run such that…"},
	{"prefix-sum", "prefix sums (+ hash map)", "range sums; count subarrays summing to k when negatives are allowed"},
	{"kadane", "Kadane / best run ending here", "the best contiguous subarray sum or product"},
	{"divide-conquer", "divide and conquer", "split in half, solve each side, combine"},
	{"partition", "partitioning (Dutch flag, quickselect)", "group around a pivot in one pass; kth element without a full sort"},
	{"binary-search", "binary search on a sorted array", "sorted input: find a target, an insert position, first/last"},
	{"bs-answer", "binary search on the answer", "minimise/maximise X where \"does X work?\" flips once"},
	{"math", "number theory / fast power", "gcd, primes up to n, x^n in log time"},
	{"tree-dfs", "tree DFS / traversal", "visit every node; build the answer from both subtrees"},
	{"tree-bfs", "tree BFS / level order", "level by level, by depth, nearest first"},
	{"bst-bounds", "BST with bounds", "BST validity or range queries: every ancestor bounds a node"},
	{"stack", "stack (matching / undo)", "brackets, nesting, the most recent unmatched item, expression evaluation"},
	{"mono-stack", "monotonic stack", "next greater/smaller element, days until warmer"},
	{"dp", "dynamic programming", "count ways / min cost / max value where choices overlap"},
	{"grid-dfs", "grid DFS/BFS (flood fill)", "connected regions of cells, islands, fill"},
	{"bfs-shortest", "BFS shortest path", "fewest steps, moves or edges; every step costs the same"},
	{"graph-traversal", "graph DFS/BFS on an adjacency list", "reachability, components, visit order"},
	{"topo-sort", "topological sort", "prerequisites, dependencies, an order, a cycle in a directed graph"},
	{"dijkstra", "Dijkstra", "cheapest/shortest path with non-negative weights"},

	// Patterns with no reflex drill yet. LeetCode cards that use them show up
	// in `--gaps` as the next drill file to write.
	{"heap", "heap / top-k", "k largest/smallest/most frequent, merge k sorted, running median"},
	{"intervals", "intervals (sort + sweep)", "overlapping ranges, meetings, merge/insert intervals"},
	{"greedy", "greedy", "one locally best choice is provably safe: reach, jumps, scheduling"},
	{"backtracking", "backtracking", "all subsets/permutations/combinations, generate every valid…"},
	{"union-find", "union-find", "merging groups, dynamic connectivity, a redundant edge"},
	{"fast-slow", "fast/slow pointers", "linked-list cycle, middle node, nth from the end"},
	{"linked-list", "linked-list rewiring", "reverse, merge or reorder nodes in place"},
	{"bits", "bit manipulation", "XOR pairs cancel, powers of two, counting bits"},
	{"trie", "trie", "prefix search over many words"},
}

// fnPattern names the pattern each reflex drill function practises. Every
// catalog function must have one (primaries_test.go checks this in study_play).
var fnPattern = map[string]string{
	"reverseInPlace": "two-pointers", "rotateRight": "two-pointers", "runningSum": "prefix-sum",
	"subarraySumK": "prefix-sum", "productExceptSelf": "prefix-sum", "maxSubarraySum": "kadane",
	"mergeSort": "divide-conquer", "quickSort": "divide-conquer", "sieve": "math",

	"twoSum": "hash-lookup", "containsDuplicate": "hash-lookup", "frequencyMap": "hash-lookup",
	"firstUniqueChar": "hash-lookup", "groupAnagrams": "hash-lookup", "singleNumber": "bits",

	"removeDuplicates": "read-write", "moveZeroes": "read-write", "maxArea": "two-pointers",
	"isPalindrome": "two-pointers", "maxSumSubarrayK": "fixed-window",
	"longestUniqueSubstring": "variable-window", "dutchFlag": "partition",

	"binarySearch": "binary-search", "searchInsert": "binary-search", "findMinRotated": "binary-search",
	"minEatingSpeed": "bs-answer", "quickSelect": "partition", "fastPow": "math", "gcd": "math",

	"inorderTraversal": "tree-dfs", "preorderTraversal": "tree-dfs", "postorderTraversal": "tree-dfs",
	"levelOrderTraversal": "tree-bfs", "maxDepth": "tree-dfs", "isValidBST": "bst-bounds",
	"isValidParentheses": "stack", "dailyTemperatures": "mono-stack",

	"climbStairs": "dp", "minCostClimbingStairs": "dp", "rob": "dp", "coinChange": "dp",

	"numIslands": "grid-dfs", "floodFill": "grid-dfs", "shortestPathGrid": "bfs-shortest",
	"canFinish": "topo-sort", "dfs": "graph-traversal", "bfs": "graph-traversal",
	"bfsShortestPath": "bfs-shortest", "topoSort": "topo-sort", "dijkstra": "dijkstra",
}

func patternByID(id string) (pattern, bool) {
	for _, p := range patterns {
		if p.id == id {
			return p, true
		}
	}
	return pattern{}, false
}

func patternIDs() []string {
	ids := make([]string, len(patterns))
	for i, p := range patterns {
		ids[i] = p.id
	}
	return ids
}

// resolvePattern maps typed input to a pattern id: an exact id, or a prefix
// that matches exactly one id ("var" -> "variable-window"). ok is false for
// unknown or ambiguous input.
func resolvePattern(in string) (string, bool) {
	in = strings.ToLower(strings.TrimSpace(in))
	if in == "" {
		return "", false
	}
	var hits []string
	for _, p := range patterns {
		if p.id == in {
			return p.id, true
		}
		if strings.HasPrefix(p.id, in) {
			hits = append(hits, p.id)
		}
	}
	if len(hits) == 1 {
		return hits[0], true
	}
	return "", false
}

// Spaced repetition. A clean pass moves an item up one box, so it comes back
// later; a miss drops it to box 0, so it comes back tomorrow.
var reviewIntervals = []int{1, 3, 7, 14, 30, 60}

const dateLayout = "2006-01-02"

func dateOf(t time.Time) string { return t.Format(dateLayout) }

// nextReview returns the new box and due date after a review. partial keeps
// the box but still waits that box's interval.
func nextReview(box int, passed, partial bool, now time.Time) (int, string) {
	switch {
	case passed:
		box++
	case partial:
	default:
		box = 0
	}
	if box < 0 {
		box = 0
	}
	if box >= len(reviewIntervals) {
		box = len(reviewIntervals) - 1
	}
	return box, dateOf(now.AddDate(0, 0, reviewIntervals[box]))
}

// daysBetween counts calendar days from a to b (b later is positive).
func daysBetween(a, b time.Time) int {
	ya, ma, da := a.Date()
	yb, mb, db := b.Date()
	ta := time.Date(ya, ma, da, 0, 0, 0, 0, time.UTC)
	tb := time.Date(yb, mb, db, 0, 0, 0, 0, time.UTC)
	return int(tb.Sub(ta).Hours() / 24)
}

// card is one solved LeetCode problem, written by `study_leetcode --log`.
// signal and insight are the parts worth remembering, not the code.
type card struct {
	Slug       string `json:"slug"`
	Title      string `json:"title"`
	Pattern    string `json:"pattern"`
	Signal     string `json:"signal"`
	Insight    string `json:"insight"`
	Guess      string `json:"guess"`      // pattern named before solving, on the last review
	Recognized bool   `json:"recognized"` // guess matched pattern on the last review
	Solved     string `json:"solved"`     // clean | hints | solution, on the last review
	Box        int    `json:"box"`
	Due        string `json:"due"`
	Logged     string `json:"logged"`
	Reviewed   string `json:"reviewed"`
	Reviews    int    `json:"reviews"`
	Misses     int    `json:"misses"`
}

type cardFile struct {
	Cards map[string]card `json:"cards"`
}

func cardsPath(repoRoot string) string {
	return filepath.Join(repoRoot, "drills", "leetcode", "cards.json")
}

func loadCards(repoRoot string) cardFile {
	f := cardFile{Cards: map[string]card{}}
	data, err := os.ReadFile(cardsPath(repoRoot))
	if err != nil {
		return f
	}
	if err := json.Unmarshal(data, &f); err != nil || f.Cards == nil {
		return cardFile{Cards: map[string]card{}}
	}
	return f
}

func saveCards(repoRoot string, f cardFile) error {
	path := cardsPath(repoRoot)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0o644)
}

// dueCards returns cards due on or before now, most overdue first.
func dueCards(f cardFile, now time.Time) []card {
	today := dateOf(now)
	var out []card
	for _, c := range f.Cards {
		if c.Due != "" && c.Due <= today {
			out = append(out, c)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Due != out[j].Due {
			return out[i].Due < out[j].Due
		}
		return out[i].Slug < out[j].Slug
	})
	return out
}
