package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
)

// Recognition round (--recognize)
//
// The drills tell you the pattern before you start: the weekday, the file name
// and the function name all give it away. This round does the step an
// interview actually asks for. It shows problem statements from every topic
// with no label, you name the pattern, and then it shows the wording that
// should have tipped you off and the look-alike pattern it is easy to confuse
// it with. Misses come back in the next rounds until you get them right.

type recogQ struct {
	id        string
	statement string
	pattern   string   // the expected pattern id
	alts      []string // other pattern ids that are also a correct answer
	why       string   // the signal in the wording
	lookalike string   // a pattern id it is easy to confuse it with
	notIt     string   // why the look-alike is the wrong pick
}

const (
	recognizeCount   = 5
	recognizeMaxRedo = 2
)

var recogBank = []recogQ{
	{"two-sum", "Given an unsorted integer array and a target, return the indices of the two numbers that add up to the target.",
		"hash-lookup", nil, "unsorted + a pair with a known sum: for each x, has target-x been seen?",
		"two-pointers", "ends-inward needs sorted input, and sorting loses the original indices"},
	{"two-sum-sorted", "Given an array sorted in non-decreasing order and a target, return the positions of two numbers that add up to the target using O(1) extra space.",
		"two-pointers", nil, "sorted + a pair + O(1) space: too small moves left up, too big moves right down",
		"hash-lookup", "a map works but costs O(n) space; sorted order is what lets the pointers move"},
	{"subarray-sum-k", "Return the number of contiguous subarrays whose sum equals k. Values may be negative.",
		"prefix-sum", nil, "count subarrays with an exact sum, negatives allowed: sum(i..j) = prefix[j+1]-prefix[i]",
		"variable-window", "with negatives, growing the window can lower the sum, so shrinking is never safe"},
	{"min-size-subarray", "Given an array of positive integers and a target, return the minimal length of a contiguous subarray whose sum is at least the target.",
		"variable-window", nil, "positive values + shortest contiguous run meeting a condition",
		"prefix-sum", "prefix sums work too, but with all-positive values the window is O(n) and O(1) space"},
	{"max-avg-k", "Return the maximum average value of any contiguous subarray of length k.",
		"fixed-window", nil, "a contiguous run of exactly k",
		"variable-window", "the length is given, so the window never needs to shrink on a condition"},
	{"longest-two-distinct", "Find the length of the longest substring that contains at most two distinct characters.",
		"variable-window", nil, "longest contiguous run such that a condition holds",
		"fixed-window", "no length is given; the left edge moves only when the condition breaks"},
	{"koko", "There are piles of bananas and h hours. Each hour you eat up to s bananas from one pile. Find the minimum integer s that lets you finish in time.",
		"bs-answer", nil, "minimise s where \"can I finish at speed s?\" flips from false to true once",
		"binary-search", "the input is not sorted; you search the range of possible speeds"},
	{"ship-capacity", "Packages must ship in the given order within D days. Find the least ship capacity that makes this possible.",
		"bs-answer", nil, "least capacity such that a feasibility check passes; feasibility is monotonic",
		"greedy", "the greedy day-by-day loading is only the check for one capacity"},
	{"search-insert", "Given a sorted array of distinct integers and a target, return the index where the target is, or where it would be inserted.",
		"binary-search", nil, "sorted input + a position: first index with nums[i] >= target",
		"bs-answer", "you search positions in the array itself, not a range of candidate answers"},
	{"rotated-min", "A sorted array was rotated an unknown number of times. Find its minimum element in O(log n).",
		"binary-search", nil, "sorted (rotated) + O(log n): compare mid with the right end to keep the half holding the drop",
		"partition", "nothing is rearranged; you only discard halves"},
	{"daily-temps", "For each day's temperature, return how many days you must wait until a warmer temperature, or 0 if none comes.",
		"mono-stack", nil, "the next greater element to the right, for every index",
		"stack", "a plain stack matches pairs; this needs a stack kept in decreasing order so each warmer day resolves the waiting ones"},
	{"valid-brackets", "Given a string of '(', ')', '[', ']', '{', '}', decide whether every closing bracket matches the most recent unmatched opener.",
		"stack", nil, "\"most recent unmatched\" is last-in, first-out",
		"mono-stack", "nothing is kept in sorted order; you just match the top"},
	{"rpn", "Evaluate an arithmetic expression written in Reverse Polish Notation.",
		"stack", nil, "operands wait until their operator arrives: push numbers, pop two on an operator",
		"mono-stack", "no ordering on the stack, only last-in first-out"},
	{"next-greater-circular", "For every element of a circular array, find the next greater element, or -1.",
		"mono-stack", nil, "next greater element; circular means walk the array twice",
		"two-pointers", "pointers cannot answer \"next greater\" for every index in one pass"},
	{"climb-stairs", "Count the distinct ways to reach the top of n stairs taking 1 or 2 steps at a time.",
		"dp", nil, "count ways, and each step's count reuses smaller counts",
		"backtracking", "listing every path is exponential; you only need the count"},
	{"coin-change", "Given coin denominations and an amount, return the fewest coins that make up the amount, or -1.",
		"dp", nil, "fewest items to reach an exact total, items reusable",
		"greedy", "largest-coin-first fails: coins [1,3,4], amount 6 is 3+3, not 4+1+1"},
	{"house-robber", "Houses in a row hold money. Adjacent houses cannot both be robbed. Maximise the total.",
		"dp", nil, "max value with a no-two-adjacent choice at each position",
		"greedy", "taking the biggest house first can block two better neighbours"},
	{"subsets", "Return every subset of a set of distinct integers.",
		"backtracking", nil, "\"return every\" valid combination",
		"dp", "you must output all 2^n subsets, not count or optimise them"},
	{"top-k-frequent", "Return the k most frequent elements of an array.",
		"heap", []string{"hash-lookup"}, "k most frequent: count with a map, then keep the top k in a size-k min-heap",
		"partition", "quickselect also works on the counts, but the heap is the standard answer"},
	{"merge-intervals", "Given a list of intervals, merge all overlapping ones.",
		"intervals", nil, "overlapping ranges: sort by start, extend the last merged interval or start a new one",
		"greedy", "it is greedy in spirit, but the move to remember is sort by start then sweep"},
	{"num-islands", "Given a grid of '1' (land) and '0' (water), count the islands.",
		"grid-dfs", []string{"union-find"}, "connected regions of cells",
		"bfs-shortest", "no distances are asked; you only need to visit each region once"},
	{"rotting-oranges", "Each minute, rotten oranges rot their fresh 4-neighbours. Return the minutes until no fresh orange is left, or -1.",
		"bfs-shortest", nil, "spreading one step per minute from many sources at once: multi-source BFS by levels",
		"grid-dfs", "DFS reaches every cell but not in minute order"},
	{"shortest-binary-matrix", "Return the length of the shortest clear path from the top-left to the bottom-right of a binary matrix, moving in 8 directions.",
		"bfs-shortest", nil, "fewest moves where every move costs the same",
		"dijkstra", "all edges weigh 1, so plain BFS is already shortest"},
	{"network-delay", "A network has directed edges with travel times. How long until a signal sent from node k reaches every node?",
		"dijkstra", nil, "shortest paths with different non-negative weights",
		"bfs-shortest", "BFS counts edges, not travel time"},
	{"course-schedule", "There are n courses and prerequisite pairs [a, b] meaning b comes before a. Can you finish every course?",
		"topo-sort", nil, "prerequisites form a directed graph; finishing all of them means there is no cycle",
		"graph-traversal", "plain reachability does not detect a directed cycle"},
	{"valid-tree", "Given n nodes and an undirected edge list, decide whether the edges form a valid tree.",
		"union-find", []string{"graph-traversal"}, "undirected: a tree has n-1 edges, no cycle, one component",
		"topo-sort", "topological order is for directed graphs"},
	{"linked-cycle", "Decide whether a singly linked list has a cycle, using O(1) extra memory.",
		"fast-slow", nil, "cycle + O(1) memory: a fast pointer laps a slow one inside a loop",
		"hash-lookup", "a set of seen nodes works but uses O(n) memory"},
	{"single-number", "Every element appears twice except one. Find that one in linear time with O(1) extra space.",
		"bits", nil, "pairs cancel and O(1) space: XOR",
		"hash-lookup", "counting works but needs O(n) space"},
	{"validate-bst", "Decide whether a binary tree is a valid binary search tree.",
		"bst-bounds", nil, "BST validity: each node must fit every ancestor's bound, not just its parent's",
		"tree-dfs", "a plain parent-child comparison misses violations deeper down"},
	{"level-order", "Return a binary tree's node values level by level, left to right.",
		"tree-bfs", nil, "level by level",
		"tree-dfs", "DFS can do it with a depth index, but a queue gives the order directly"},
	{"diameter", "Return the number of edges on the longest path between any two nodes of a binary tree.",
		"tree-dfs", nil, "the answer at a node combines the heights of both subtrees",
		"tree-bfs", "levels do not tell you path lengths through a node"},
	{"sort-colors", "Sort an array of 0s, 1s and 2s in place in one pass.",
		"partition", nil, "three groups, one pass, in place: low/mid/high pointers",
		"divide-conquer", "a full sort is O(n log n) and more than one pass"},
	{"kth-largest", "Find the kth largest element of an unsorted array without sorting the whole array.",
		"partition", []string{"heap"}, "kth element without a full sort: quickselect (or a size-k min-heap)",
		"binary-search", "the array is not sorted"},
	{"move-zeroes", "Move every 0 to the end of the array in place, keeping the relative order of the other elements.",
		"read-write", nil, "in place + keep order: a write pointer for the next kept element",
		"two-pointers", "swapping from both ends breaks the order"},
	{"container-water", "Pick two vertical lines that, together with the x-axis, hold the most water.",
		"two-pointers", nil, "a pair whose score depends on both ends: start wide, move the shorter wall",
		"dp", "there is no overlapping subproblem; the pointers rule out pairs safely"},
	{"max-subarray", "Return the largest sum of any non-empty contiguous subarray.",
		"kadane", []string{"dp"}, "the best contiguous sum: best run ending here = max(x, previous + x)",
		"prefix-sum", "prefix sums can do it, but Kadane needs one pass and O(1) space"},
	{"autocomplete", "Store many words and answer many queries for all words that start with a given prefix.",
		"trie", nil, "many prefix queries over a word list",
		"hash-lookup", "a map finds whole words, not every word under a prefix"},
	{"jump-game", "Each element is your maximum jump length from that index. Can you reach the last index?",
		"greedy", []string{"dp"}, "track the furthest reachable index; one locally best choice is safe",
		"dp", "DP works but is O(n^2); the reach greedy is O(n)"},
	{"reverse-list", "Reverse a singly linked list in place.",
		"linked-list", nil, "rewire next pointers in place: prev, cur, next",
		"stack", "a stack works but uses O(n) memory"},
	{"count-primes", "Count the prime numbers less than n.",
		"math", nil, "all primes up to n: sieve of Eratosthenes",
		"dp", "nothing is optimised; you cross out multiples"},
}

// pickRecognition chooses today's round: up to recognizeMaxRedo recently
// missed questions first, then questions for patterns not already in the
// round, rotating by date so each day differs.
func pickRecognition(bank []recogQ, misses map[string]skillRecord, now time.Time, n int) []recogQ {
	byID := map[string]recogQ{}
	for _, q := range bank {
		byID[q.id] = q
	}
	var redo []string
	for id, rec := range misses {
		t, err := time.Parse(dateLayout, rec.LastMiss)
		if _, ok := byID[id]; !ok || err != nil || rec.Misses <= 0 || daysBetween(t, now) > skillMissWindow {
			continue
		}
		redo = append(redo, id)
	}
	sort.Slice(redo, func(i, j int) bool {
		a, b := misses[redo[i]], misses[redo[j]]
		if a.Misses != b.Misses {
			return a.Misses > b.Misses
		}
		return redo[i] < redo[j]
	})

	var out []recogQ
	used := map[string]bool{}
	usedPattern := map[string]bool{}
	take := func(q recogQ) {
		out = append(out, q)
		used[q.id] = true
		usedPattern[q.pattern] = true
	}
	for _, id := range redo {
		if len(out) >= recognizeMaxRedo || len(out) >= n {
			break
		}
		take(byID[id])
	}
	if len(bank) == 0 {
		return out
	}
	// Walk the bank with a stride coprime to its size, so picks spread across
	// topics and every question is reachable.
	stride := 7
	for gcd(stride, len(bank)) != 1 {
		stride++
	}
	start := dayNumber(now) * n
	for pass := 0; pass < 2 && len(out) < n; pass++ {
		for i := 0; i < len(bank) && len(out) < n; i++ {
			q := bank[(start+i*stride)%len(bank)]
			if used[q.id] || (pass == 0 && usedPattern[q.pattern]) {
				continue
			}
			take(q)
		}
	}
	return out
}

// cardQuestions turns LeetCode cards (bin/study_leetcode --log) into
// recognition questions, so problems you solved come back here too.
func cardQuestions(f cardFile) []recogQ {
	var out []recogQ
	for _, c := range f.Cards {
		if _, ok := patternByID(c.Pattern); !ok {
			continue
		}
		q := recogQ{
			id:        "lc:" + c.Slug,
			statement: fmt.Sprintf("LeetCode: %s (you solved it on %s) — which pattern?", c.Title, c.Logged),
			pattern:   c.Pattern,
			why:       c.Signal,
		}
		if c.Insight != "" {
			q.notIt = "your insight: " + c.Insight
		}
		out = append(out, q)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].id < out[j].id })
	return out
}

func (q recogQ) accepts(id string) bool {
	if id == q.pattern {
		return true
	}
	for _, a := range q.alts {
		if id == a {
			return true
		}
	}
	return false
}

// runRecognition asks each question on out, reads one answer per line from in,
// and returns the ids answered right and wrong. "?" or an empty line counts as
// a miss; an unknown pattern id is asked again. At EOF the remaining questions
// are skipped and not graded.
func runRecognition(qs []recogQ, in io.Reader, out io.Writer) (right, wrong []string) {
	fmt.Fprintln(out, "── RECOGNITION — name the pattern before anything else ─")
	fmt.Fprintln(out, "  Answer with a pattern id (a unique prefix is enough), or ? if you do not know.")
	fmt.Fprintln(out)
	fmt.Fprint(out, formatInColumns(patternIDs(), "  ", longestPatternID(), 80))
	sc := bufio.NewScanner(in)
	for i, q := range qs {
		fmt.Fprintf(out, "\n%d/%d  %s\n", i+1, len(qs), q.statement)
		var got string
		answered := false
		for tries := 0; tries < 3; tries++ {
			fmt.Fprint(out, "  pattern> ")
			if !sc.Scan() {
				fmt.Fprintln(out, "\n  (no more input — the rest is not graded)")
				return right, wrong
			}
			line := strings.TrimSpace(sc.Text())
			if line == "" || line == "?" {
				answered = true
				break
			}
			if id, ok := resolvePattern(line); ok {
				got, answered = id, true
				break
			}
			fmt.Fprintf(out, "  %q is not a pattern id (or matches several); try again\n", line)
		}
		if !answered {
			got = ""
		}
		if q.accepts(got) {
			fmt.Fprintf(out, "  ✓ %s\n", q.pattern)
			right = append(right, q.id)
		} else {
			fmt.Fprintf(out, "  ✗ answer: %s\n", q.pattern)
			wrong = append(wrong, q.id)
		}
		if q.why != "" {
			fmt.Fprintf(out, "    signal: %s\n", q.why)
		}
		if q.lookalike != "" {
			fmt.Fprintf(out, "    not %s: %s\n", q.lookalike, q.notIt)
		} else if q.notIt != "" {
			fmt.Fprintf(out, "    %s\n", q.notIt)
		}
	}
	fmt.Fprintf(out, "\n%d/%d recognised.", len(right), len(qs))
	if len(wrong) > 0 {
		fmt.Fprint(out, " Misses come back in the next rounds.")
	}
	fmt.Fprintln(out)
	return right, wrong
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func longestPatternID() int {
	n := 0
	for _, p := range patterns {
		if len(p.id) > n {
			n = len(p.id)
		}
	}
	return n
}

// recordRecognition stores misses like small skills: a miss adds one, a right
// answer removes one, and zero drops out of the log.
func recordRecognition(log *drillLog, right, wrong []string, now time.Time) {
	if log.Recognition == nil {
		log.Recognition = map[string]skillRecord{}
	}
	for _, id := range wrong {
		rec := log.Recognition[id]
		rec.Misses++
		rec.LastMiss = dateOf(now)
		log.Recognition[id] = rec
	}
	for _, id := range right {
		rec, ok := log.Recognition[id]
		if !ok {
			continue
		}
		rec.Misses--
		if rec.Misses <= 0 {
			delete(log.Recognition, id)
		} else {
			log.Recognition[id] = rec
		}
	}
}

// recognize runs one round against stdin/stdout and saves the result.
func recognize(repoRoot string, in io.Reader, out io.Writer) error {
	now := time.Now()
	log := loadLog(repoRoot)
	bank := append(append([]recogQ{}, recogBank...), cardQuestions(loadCards(repoRoot))...)
	qs := pickRecognition(bank, log.Recognition, now, recognizeCount)
	right, wrong := runRecognition(qs, in, out)
	if len(right)+len(wrong) == 0 {
		return nil
	}
	recordRecognition(&log, right, wrong, now)
	return saveLog(repoRoot, log)
}
