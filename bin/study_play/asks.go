package main

import (
	"fmt"
	"time"
)

type askPrompt struct {
	statement string
	hints     []string
}

// skillQ is one small-skill question: index math, range models, invariants,
// pointer reasons, state meaning, complexity. It sits beside the daily ask so
// the mechanics under today's pattern are rehearsed before the code is.
type skillQ struct {
	skill  string
	q      string
	answer string
}

// revealSkillAnswers is set by --show so the answers print under each question.
var revealSkillAnswers bool

var dailyAsks = map[string]askPrompt{
	"Monday": {
		statement: "Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.",
		hints: []string{
			"Ask: Move last k elements to the front",
			"Input: array, k (k may be larger than n)",
			"Output: rotated array (in-place or new?)",
			"Edge: empty, k=0, k % n",
			"Pattern: reverse sections or modulo indexing",
		},
	},
	"Tuesday": {
		statement: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.",
		hints: []string{
			"Ask: Find two indices whose values sum to target",
			"Input: unsorted ints, one valid pair guaranteed",
			"Output: two indices (not values)",
			"Edge: negatives, duplicate values",
			"Pattern: hash map + complement",
		},
	},
	"Wednesday": {
		statement: "You are given an integer array height of length n. There are n vertical lines. Find two lines that together with the x-axis form a container that holds the most water.",
		hints: []string{
			"Ask: Maximize area = width × min(height[i], height[j])",
			"Input: heights array",
			"Output: max area (integer)",
			"Edge: two elements, all same height",
			"Pattern: two pointers L/R",
		},
	},
	"Thursday": {
		statement: "Given a sorted array of distinct integers and a target value, return the index if found. If not, return the index where it would be inserted.",
		hints: []string{
			"Ask: Lower bound / insertion position",
			"Input: sorted distinct array, target",
			"Output: index (0..n)",
			"Edge: insert before all, after all",
			"Pattern: binary search, first index ≥ target",
		},
	},
	"Friday": {
		statement: "Given a string s containing just '(', ')', '{', '}', '[' and ']', determine if the input string is valid.",
		hints: []string{
			"Ask: Every closing bracket matches most recent unmatched opener",
			"Input: bracket string",
			"Output: boolean",
			"Edge: empty, single opener, wrong type close",
			"Pattern: stack",
		},
	},
	"Saturday": {
		statement: "You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay cost[i], you can climb one or two steps. Return the minimum cost to reach the top.",
		hints: []string{
			"Ask: Min total cost to reach index n (past last step)",
			"Input: cost per step",
			"Output: min total cost",
			"Edge: n=1, n=2",
			"Pattern: 1D DP — dp[i] = cost[i] + min(dp[i-1], dp[i-2])",
		},
	},
	"Sunday": {
		statement: "Given an m×n 2D binary grid which represents a map of '1's (land) and '0's (water), return the number of islands.",
		hints: []string{
			"Ask: Count connected components of land",
			"Input: grid of '0' and '1'",
			"Output: island count",
			"Edge: all water, all land, single cell",
			"Pattern: DFS/BFS + visited",
		},
	},
}

func printAskWarmup(day string) {
	ask, ok := dailyAsks[day]
	if !ok {
		return
	}
	fmt.Println("\n── QUESTION LITERACY (before coding) ──────────────────")
	fmt.Printf("  Problem: %s\n\n", ask.statement)
	fmt.Println("  Fill in (on paper or aloud):")
	for _, h := range ask.hints {
		fmt.Printf("    • %s\n", h)
	}
	printSkillQuestions(day)
	fmt.Println("\n  Full asks pack: bin/study_play/_support/asks/README.md")
}

// Small-skill questions
//
// Writing a function from memory and deriving one are different skills. These
// are the small moves underneath the patterns (see "Mechanics" in
// doc/write/STUDY_PLAN.md), drilled as short questions with one checkable
// answer. Each day shows two blocks:
//
//	focus  — skills that feed today's weekday drill, rotating week by week
//	cross  — one question from each of several other skills, walking the whole
//	         bank day by day so no topic is left out
//
// Both are derived from the date, so a day's set is stable and tomorrow's
// differs. Every question in the bank comes up in rotation.
const (
	skillFocusSkills   = 2 // skills drawn from the weekday's focus list
	skillFocusPerSkill = 2 // questions per focus skill
	skillCrossCount    = 4 // redo + cross-topic questions together
	skillMaxRedo       = 2 // at most this many questions come from missed skills
	skillMissWindow    = 14
)

// skillOrder is the dependency order the cross-topic walk follows: mechanics
// first, then patterns, structures, recursion, graphs, DP, reasoning.
var skillOrder = []string{
	"index",
	"range",
	"coord",
	"mapping",
	"state",
	"invariant",
	"pointer",
	"prefix",
	"freq",
	"bsearch",
	"mono",
	"stack",
	"list",
	"tree",
	"heap",
	"recur",
	"graph",
	"dp",
	"greedy",
	"complexity",
	"bits",
}

// skillFocus lines the small skills up with the weekday drill they warm up.
var skillFocus = map[string][]string{
	"Monday":    {"index", "range", "prefix", "coord", "mapping", "bits"},
	"Tuesday":   {"freq", "state", "mapping", "complexity", "bits"},
	"Wednesday": {"pointer", "invariant", "state", "range", "complexity"},
	"Thursday":  {"bsearch", "mono", "invariant", "range", "bits"},
	"Friday":    {"stack", "tree", "list", "heap", "mono"},
	"Saturday":  {"dp", "recur", "greedy", "state", "complexity"},
	"Sunday":    {"graph", "coord", "recur", "stack", "greedy"},
}

var skillBank = []skillQ{
	{"index", "n=10. What is the mirror of index i=3 when reversing? Give the formula too.",
		"n-1-i = 6"},
	{"index", "n=7. A cursor at index 5 moves 4 steps right and wraps around the end. Where does it land?",
		"(5+4)%7 = 2"},
	{"index", "Rotate right by k=12 on an array of n=5. What is the effective shift, and where does the element at index 4 end up?",
		"k%n = 2; (4+2)%5 = 1"},
	{"index", "A loop reads nums[i+1]. What must the loop condition be, and why is i<n wrong?",
		"i+1<n (i<=n-2); at i=n-1 the read is nums[n], out of range"},
	{"index", "Which index is the last valid one, which is the first invalid one, and what do they equal for n=0?",
		"last = n-1, first invalid = n; for n=0 last is -1, so nothing is valid"},
	{"range", "How many elements are in the inclusive window [3,7]? In the half-open [3,7)?",
		"5 (r-l+1); 4 (r-l)"},
	{"range", "n=10 and windows of size k=3. How many windows exist and what is the last window's start index?",
		"n-k+1 = 8 windows; last start is n-k = 7"},
	{"range", "Binary search over the half-open [lo,hi) with hi=n. Give the loop condition and both updates.",
		"lo<hi; hi=mid (not mid-1); lo=mid+1"},
	{"range", "With window [left,right] inclusive, right just advanced. Write the window length. If left==right what is it?",
		"right-left+1; 1"},
	{"range", "Go: nums[2:9] — what is its length, and is index 9 part of it?",
		"7 elements (9-2); index 9 is excluded"},
	{"range", "How many pairs (i<j) can you pick from n=6 elements?",
		"n(n-1)/2 = 15"},
	{"coord", "A grid has 3 rows and 4 columns. What is the flat index of cell (r=2,c=1)?",
		"r*cols+c = 2*4+1 = 9"},
	{"coord", "Same 3x4 grid. Which (r,c) is flat index 10?",
		"r=10/4=2, c=10%4=2"},
	{"coord", "3x4 grid, cell (0,3). How many of its 4-neighbours are inside the grid, and which bound check rejects the others?",
		"2 — (1,3) and (0,2); r-1<0 and c+1==cols fail"},
	{"coord", "Rotate an n x n matrix 90 degrees clockwise. Where does cell (r,c) go? Test it on (0,1) with n=3.",
		"(c, n-1-r); (0,1) -> (1,2)"},
	{"coord", "Which quantity is constant along an anti-diagonal and which along a main diagonal?",
		"anti-diagonal: r+c; main diagonal: r-c"},
	{"mapping", "Values are in 1..n. At which index does value 4 belong, and how do you find a duplicate by that rule?",
		"index 3 (v-1); a value whose home slot already holds that value is a duplicate"},
	{"mapping", "Extract and drop the last digit of x=472, then build the reverse of 123 digit by digit.",
		"472%10=2, 472/10=47; rev=rev*10+d gives 3, 32, 321"},
	{"mapping", "Map a lowercase letter to a 0..25 slot. What is the slot of 'e'?",
		"'e'-'a' = 4"},
	{"mapping", "nums=[0,1,3], n=3: one number in 0..n is missing. Find it with a formula.",
		"n(n+1)/2 - sum = 6 - 4 = 2"},
	{"mapping", "Read nums as 'index i points to index nums[i]'. Why does a duplicate in 1..n force a cycle?",
		"two indices point to the same target, so some node has in-degree 2 — a list that loops back"},
	{"state", "In a sliding window, what does 'sum' mean? Say it as one phrase.",
		"the sum of nums[left..right] — the current window, nothing else"},
	{"state", "Two-sum with a map: what does the map hold when you reach index i, and why look up before inserting nums[i]?",
		"value -> index of elements before i; inserting first could pair an element with itself"},
	{"state", "BFS shortest path: what does dist[v] mean, and when is it written?",
		"edges from the source to v; written once, when v is first enqueued"},
	{"state", "Coin change: what does dp[a] mean, and why is dp[0]=0 while every other cell starts at 'infinity'?",
		"fewest coins making exactly a; zero coins make 0, every other amount is unproven"},
	{"state", "Valid parentheses: what does the stack contain at any moment?",
		"the openers not yet matched, oldest at the bottom"},
	{"state", "Longest substring without repeats: left, right, best, last-seen. Give each one's meaning.",
		"left = start of the current valid window; right = its end; best = longest valid window so far; last-seen = most recent index of each char"},
	{"invariant", "Longest substring without repeating characters: state the invariant on [left,right].",
		"every character in [left,right] is unique"},
	{"invariant", "Exact binary search with lo<=hi. What is the invariant on [lo,hi]?",
		"if the target exists, it is inside [lo,hi]"},
	{"invariant", "Dutch flag with low, mid, high: say what each of the four regions holds.",
		"[0,low)=0s, [low,mid)=1s, [mid,high]=unknown, (high,n)=2s"},
	{"invariant", "In the Dutch flag, why does swapping with high NOT advance mid?",
		"the value that arrives from high is unknown and still needs inspecting"},
	{"invariant", "Insertion sort: what is true after outer iteration i finishes?",
		"nums[0..i] is sorted"},
	{"invariant", "Kadane: what does cur mean at index i, and what does best mean?",
		"cur = best sum of a subarray ending exactly at i; best = max cur seen so far"},
	{"pointer", "Sorted array two-sum, sum > target. Which pointer moves and what did you just learn?",
		"right--: nums[right] with any left' >= left gives an even larger sum, so right is useless"},
	{"pointer", "Container with most water: why move the shorter wall inward?",
		"width shrinks either way, so only a taller minimum can help; moving the taller wall cannot raise the minimum"},
	{"pointer", "removeDuplicates read/write pointers: what does the write pointer represent?",
		"the next free slot; nums[:w] is the deduplicated prefix"},
	{"pointer", "Slow moves 1, fast moves 2. On a 5-node list where is slow when fast stops? On a 6-node list?",
		"index 2 (the middle); index 3 (second middle)"},
	{"pointer", "Variable window: when does left move, and can it ever move backward?",
		"only while the window is invalid; never backward"},
	{"pointer", "Merging two sorted arrays: which pointer advances, and what do you do when one array runs out?",
		"the one whose element was taken; copy the remaining tail of the other"},
	{"prefix", "nums=[2,4,6,8]. Write the prefix array with a leading 0, then the sum of nums[1..2] inclusive.",
		"[0,2,6,12,20]; prefix[3]-prefix[1] = 10"},
	{"prefix", "Give the formula for the sum of inclusive [l,r] using a length n+1 prefix array.",
		"prefix[r+1]-prefix[l]"},
	{"prefix", "Subarray-sum-equals-k with a map of prefixes: why seed the map with {0:1}?",
		"so subarrays starting at index 0 are counted — they equal prefix minus the empty prefix"},
	{"prefix", "nums=[1,2,3,4]. Product of everything except index 1, using prefix and suffix products.",
		"prefixProd(before 1)=1 times suffixProd(after 1)=3*4=12"},
	{"prefix", "Add 5 to every element of [1,3] in a length-5 zero array using a difference array. What are the two writes and the final array?",
		"diff[1]+=5, diff[4]-=5; prefix-sum gives [0,5,5,5,0]"},
	{"freq", "Anagram check with a 26-slot array: what must be true at the end?",
		"every slot is 0 (add one string, subtract the other)"},
	{"freq", "Why is the sum of character codes a bad group key for anagrams? Name a better key.",
		"'ad' and 'bc' both sum to 197; use the sorted string or a 26-count signature"},
	{"freq", "A value appears c=5 times. How many equal pairs (i<j) does it contribute?",
		"c(c-1)/2 = 10"},
	{"freq", "Sliding a fixed window: how many map updates happen per step, and what do they cost?",
		"two — one char enters, one leaves; O(1) each"},
	{"freq", "Set or map: 'has this been seen?' versus 'where / how many times?'",
		"set; map"},
	{"freq", "Track the number of distinct values in a window without scanning the map. When does the counter change?",
		"+1 when a count goes 0->1; -1 when it goes 1->0"},
	{"bsearch", "Lower bound (first index with nums[i] >= target) on [1,2,2,2,5]. Answers for targets 2, 3, and 6.",
		"1, 4, 5 (5 = n means 'insert at the end')"},
	{"bsearch", "Upper bound is the first index with nums[i] > target. On [1,2,2,2,5], how do you count the 2s?",
		"upper(2)-lower(2) = 4-1 = 3"},
	{"bsearch", "Piles [3,6,7,11], h=8 hours. Minimum eating speed? State the predicate and why it can be searched.",
		"4 (4:1+2+2+3=8 ok, 3:10 too slow); 'finishes in <=h' flips false->true once as speed grows"},
	{"bsearch", "Loop uses lo=mid on lo=3, hi=4 with mid=lo+(hi-lo)/2. What goes wrong and what is the fix?",
		"mid=3 so lo never moves — infinite loop; use the upper mid lo+(hi-lo+1)/2 whenever lo=mid"},
	{"bsearch", "Rotated sorted [4,5,6,7,0,1,2]: lo=0, hi=6. After the first mid compare, which half survives and why?",
		"mid=3, 7>nums[hi]=2, so the minimum is in (mid,hi]; lo=4"},
	{"bsearch", "Exact search on a sorted array: what are the two things mid can tell you, and which half is eliminated for each?",
		"nums[mid]<target -> discard [lo,mid]; nums[mid]>target -> discard [mid,hi]; equal -> found"},
	{"mono", "Which is monotonic in the speed s: 'can finish all piles within h' or 's is a perfect square'?",
		"the first — a bigger speed never hurts; the second flips back and forth"},
	{"mono", "Next greater element for [2,1,5,3] with a monotonic stack. Give the answers and what is popped when 5 arrives.",
		"[5,5,-1,-1]; pop 1 then 2 — both get 5"},
	{"mono", "Sliding window maximum with a deque of indices, window size 3, current index i=5. Which front index is stale?",
		"any index <= i-k = 2"},
	{"mono", "Why does the variable-window 'shortest subarray with sum >= target' break once negatives are allowed?",
		"extending the window no longer increases the sum, so shrinking is no longer safe"},
	{"mono", "A decreasing stack of values: which direction does it find, and what does an element's pop tell you?",
		"next greater; the popper is that element's next greater"},
	{"stack", "Evaluate RPN '3 4 + 2 *'. Show the stack after each token.",
		"[3] [3,4] [7] [7,2] [14] -> 14"},
	{"stack", "Daily temperatures: the stack stores indices, not values. Why? Answer for [73,74,75,71,69,72,76,73] at index 2.",
		"you need the distance between indices; 4 (the 76 at index 6)"},
	{"stack", "Level-order BFS: why capture size := len(queue) before the inner loop?",
		"the queue grows while you push children; size fixes the boundary of the current level"},
	{"stack", "Iterative preorder with a stack: push order of the two children, and why?",
		"push right then left, so left pops first"},
	{"stack", "Pick stack or queue: DFS, BFS, undo/matching brackets, shortest path in edges.",
		"stack, queue, stack, queue"},
	{"list", "Reverse a list: why save cur.next before setting cur.next=prev?",
		"otherwise the rest of the list is unreachable"},
	{"list", "When do you reach for a dummy head node?",
		"whenever the head itself can be removed or replaced — it removes the first-node special case"},
	{"list", "Remove the n-th node from the end: describe the gap trick.",
		"advance fast n steps first, then move both until fast.next is nil; slow.next is the target (start from a dummy)"},
	{"list", "Why must fast and slow meet if a cycle exists?",
		"inside the cycle the gap shrinks by 1 each step, so it hits 0"},
	{"list", "Merging two sorted lists: after the loop one list still has nodes. What do you do?",
		"tail.next = the non-empty remainder — it is already linked and sorted, no loop needed"},
	{"list", "How many .next hops reach the node at index i from head?",
		"i hops"},
	{"tree", "What does an inorder traversal of a BST produce?",
		"the values in sorted ascending order"},
	{"tree", "Preorder [3,9,20,15,7], inorder [9,3,15,20,7]. Which is the root and how big is the left subtree?",
		"root 3 (first of preorder); left size 1 (index of 3 in inorder)"},
	{"tree", "Diameter of a tree: what does the recursive call return versus record?",
		"returns height 1+max(l,r) to its parent; records l+r into the global best"},
	{"tree", "Why does isValidBST carry (lo,hi) bounds instead of comparing a node to its parent?",
		"a deep node must respect every ancestor, not just the immediate one"},
	{"tree", "Perfect binary tree with 4 levels: how many nodes? How about DFS vs BFS extra space?",
		"2^4-1 = 15; DFS O(height), BFS O(widest level)"},
	{"tree", "maxDepth: what do nil and a single node return?",
		"0 and 1 (depth counted in nodes)"},
	{"heap", "Min-heap array [1,3,2,7,4]. Index of the parent of index 4, and the children of index 1?",
		"(4-1)/2 = 1; indices 3 and 4"},
	{"heap", "Insert 0 into [1,3,2,7,4]. Show the sift-up result.",
		"[0,3,1,7,4,2]"},
	{"heap", "k-th largest with a heap: min-heap or max-heap, what size, and where is the answer?",
		"min-heap of size k; its root is the k-th largest"},
	{"heap", "Is [1,5,2,6,7,3] a valid min-heap? Is it sorted?",
		"valid (each parent <= its kids); not sorted"},
	{"heap", "Building a heap from n items: what is the cost and the first index to sift down for n=10?",
		"O(n); last internal node n/2-1 = 4, then work back to 0"},
	{"heap", "Cost of push, pop, and peek on a binary heap?",
		"O(log n), O(log n), O(1)"},
	{"recur", "How many subsets of [1,2,3], permutations of 3 items, and 2-element combinations of 4?",
		"8, 6, 6"},
	{"recur", "A backtracker appends path to results directly and the output is wrong. Why?",
		"slices share a backing array, later mutations rewrite saved results — copy the path"},
	{"recur", "Permutations: what must be restored after the recursive call returns?",
		"the used[i] flag and the last element of the path"},
	{"recur", "Combination sum: why pass a start index instead of looping from 0?",
		"so [2,3] and [3,2] are not both generated — choices only go forward"},
	{"recur", "Subsets II (sorted, duplicates): why 'i > start && nums[i]==nums[i-1]' and not 'i > 0'?",
		"a repeat is only skipped among siblings at the same depth; it is legal as the first pick at a deeper level"},
	{"recur", "Naive fib(n) is exponential; what is the state, and what does a memo change?",
		"state = n; each n is computed once, so O(n)"},
	{"recur", "A recursive call that visits n nodes in a chain: what is its space cost?",
		"O(n) call-stack depth"},
	{"graph", "Undirected edges [[0,1],[0,2],[1,2]]: write the adjacency list and the degree of node 0.",
		"0:[1,2] 1:[0,2] 2:[0,1]; degree 2"},
	{"graph", "Number of Islands: name the node, the edge, and the visited rule.",
		"node = land cell; edge = 4-neighbour land cell; mark visited when first reached"},
	{"graph", "Prerequisites [a,b] mean you must take b before a. Which way is the edge, and how do you detect a cycle?",
		"b -> a; fewer than n nodes come out of Kahn's order"},
	{"graph", "Count connected components in an undirected graph.",
		"loop over nodes; each unvisited start triggers a DFS/BFS and count++"},
	{"graph", "Adjacency matrix versus list: space and edge-lookup cost.",
		"matrix O(V^2) space, O(1) lookup; list O(V+E) space, O(degree) lookup"},
	{"graph", "Why does Dijkstra require non-negative edge weights?",
		"a popped node's distance is final; a negative edge could later undercut it"},
	{"graph", "Union-Find on 5 nodes: union(0,1), (1,2), (3,4), (0,2). How many components, and what did the last union reveal?",
		"2; 0 and 2 were already connected — that edge would close a cycle"},
	{"graph", "When do you mark a node visited in BFS: on enqueue or dequeue?",
		"on enqueue — otherwise the same node enters the queue several times"},
	{"dp", "climbStairs: write dp[i] in English, the transition, the bases, and dp[5].",
		"ways to reach step i; dp[i]=dp[i-1]+dp[i-2]; dp[0]=dp[1]=1; 8"},
	{"dp", "House robber on [2,7,9,3,1]: define dp[i] and give the final answer.",
		"best loot from the first i houses, dp[i]=max(dp[i-1], dp[i-2]+nums[i-1]); 12"},
	{"dp", "Coin change coins [1,3,4], amount 6. Result, and why greedy fails.",
		"2 (3+3); greedy takes 4+1+1 = 3 coins"},
	{"dp", "dp[i] reads only dp[i-1] and dp[i-2]. What is the space optimisation?",
		"keep two variables — O(1) space"},
	{"dp", "1D knapsack: which direction does the capacity loop run for 0/1 versus unbounded, and why?",
		"0/1 downward so each item is used once; unbounded upward so an item may repeat"},
	{"dp", "LCS table: why is it (m+1) x (n+1), and what do row 0 / column 0 mean?",
		"index = prefix length; row/col 0 are the empty prefix, the base case"},
	{"dp", "Grid paths on an m x n grid: how many states and what is the time cost?",
		"m*n states, O(1) transition each -> O(mn)"},
	{"greedy", "Max non-overlapping intervals for [1,3],[2,4],[3,5]. Which sort key and what is the answer?",
		"sort by end time; 2 ([1,3] and [3,5])"},
	{"greedy", "Merge [[1,3],[2,6],[8,10]]. Sort key and merge condition?",
		"sort by start; merge if next.start <= cur.end; result [[1,6],[8,10]]"},
	{"greedy", "Why does sorting make three-sum's two-pointer step valid?",
		"order tells you the direction: too small moves left up, too big moves right down"},
	{"greedy", "Jump game on [3,2,1,0,4]: what do you track and what is the verdict?",
		"furthest reachable index; false — it stalls at index 3"},
	{"greedy", "Meeting rooms II on [0,30],[5,10],[15,20]: how do the sorted starts and ends give the answer?",
		"walk both sorted lists, peak overlap = 2 rooms"},
	{"greedy", "Exchange argument for earliest-end-time scheduling, in one sentence?",
		"any optimal schedule's first meeting can be swapped for the earliest-ending one without conflicting with the rest"},
	{"complexity", "for i<n { for j:=i; j<n; j++ } — exact count and big-O?",
		"n(n+1)/2 iterations; O(n^2)"},
	{"complexity", "for i:=1; i<n; i*=2 with n=1024: how many iterations?",
		"10 (i = 1,2,4,...,512)"},
	{"complexity", "A while-loop inside a for-loop moves left forward only. Total cost?",
		"O(n) — left makes at most n moves across the whole run"},
	{"complexity", "Solve T(n)=2T(n/2)+n, T(n)=T(n/2)+1, T(n)=2T(n-1).",
		"O(n log n), O(log n), O(2^n)"},
	{"complexity", "A monotonic stack pushes each element once and may pop inside a loop. Total cost?",
		"O(n) — at most n pushes and n pops"},
	{"complexity", "Why is s += c in a loop O(n^2) for strings in Go, and what is the fix?",
		"strings are immutable, each concat copies; use strings.Builder"},
	{"complexity", "Appending to a doubling slice: why is it amortized O(1)?",
		"copies total 1+2+4+... < 2n across n appends"},
	{"complexity", "Rule of thumb: what complexity can n=20, n=500, and n=100000 afford?",
		"exponential/bitmask, O(n^3), O(n log n)"},
	{"bits", "Single number in [4,1,2,1,2]: what operation, why, and the answer?",
		"XOR everything; pairs cancel (a^a=0) leaving 4"},
	{"bits", "Power-of-two test and its reason, tried on 8.",
		"x>0 && x&(x-1)==0; 8&7=0"},
	{"bits", "Count the set bits of 12 using x&(x-1).",
		"12->8->0: two drops, so 2"},
	{"bits", "Is bit 2 set in 13? Write the test.",
		"yes — (13>>2)&1 = 1"},
	{"bits", "Go: -7 % 5 is what? Write a modulo that stays non-negative.",
		"-2; ((a%n)+n)%n gives 3"},
	{"bits", "gcd(48,18) by Euclid, then lcm.",
		"48%18=12, 18%12=6, 12%6=0 -> 6; lcm = 48/6*18 = 144"},
	{"bits", "Fast power 3^13: which squarings do you multiply in?",
		"13=1101b; multiply 3^1, 3^4, 3^8 (bits 0,2,3)"},
	{"bits", "Sieve of Eratosthenes up to 30: which primes do you cross out from, and where does each start?",
		"2, 3, 5 only (7^2>30); each starts at p*p"},
}

func skillQuestions(skill string) []skillQ {
	var out []skillQ
	for _, q := range skillBank {
		if q.skill == skill {
			out = append(out, q)
		}
	}
	return out
}

// pickSkillQ returns the nth question of a skill, wrapping around.
func pickSkillQ(skill string, n int) skillQ {
	qs := skillQuestions(skill)
	return qs[((n%len(qs))+len(qs))%len(qs)]
}

// dayNumber counts whole days since the Unix epoch, so the rotation advances
// by exactly one step per calendar day.
func dayNumber(now time.Time) int {
	y, m, d := now.Date()
	return int(time.Date(y, m, d, 0, 0, 0, 0, time.UTC).Unix() / 86400)
}

// skillSet returns today's focus, redo and cross-topic questions. redo holds
// questions from skills you recently marked --missed, ahead of the rotation, so
// a weak skill comes back sooner and with a different question each time. The
// three blocks always total skillFocusSkills*skillFocusPerSkill+skillCrossCount.
func skillSet(day string, now time.Time, weak []string) (focus, redo, cross []skillQ) {
	dn := dayNumber(now)
	ids := skillFocus[day]
	inFocus := map[string]bool{}
	week := dn / 7
	for i := 0; i < skillFocusSkills && len(ids) > 0; i++ {
		id := ids[(week*skillFocusSkills+i)%len(ids)]
		if inFocus[id] {
			continue
		}
		inFocus[id] = true
		for j := 0; j < skillFocusPerSkill; j++ {
			focus = append(focus, pickSkillQ(id, week*skillFocusPerSkill+j))
		}
	}

	skipped := map[string]bool{}
	for _, id := range weak {
		if inFocus[id] || len(redo) >= skillMaxRedo {
			continue
		}
		skipped[id] = true
		redo = append(redo, pickSkillQ(id, dn))
	}

	total := len(skillOrder)
	want := skillCrossCount - len(redo)
	for step := 0; len(cross) < want && step < total*2; step++ {
		pos := dn*skillCrossCount + step
		id := skillOrder[pos%total]
		if inFocus[id] || skipped[id] {
			continue
		}
		cross = append(cross, pickSkillQ(id, pos/total))
	}
	return focus, redo, cross
}

// printSkillQuestions prints today's small-skill questions. Say the answer and
// the reason aloud first; --show reveals the answers.
func printSkillQuestions(day string) {
	printSkillQuestionsAt(day, time.Now(), weakSkillsNow(time.Now()))
}

func printSkillQuestionsAt(day string, now time.Time, weak []string) {
	focus, redo, cross := skillSet(day, now, weak)
	if len(focus)+len(redo)+len(cross) == 0 {
		return
	}
	fmt.Println("\n── SMALL SKILLS (say the answer and the reason) ────────")
	n := 1
	print := func(label string, qs []skillQ) {
		if len(qs) == 0 {
			return
		}
		fmt.Printf("  %s\n", label)
		for _, q := range qs {
			fmt.Printf("  %d) [%s] %s\n", n, q.skill, q.q)
			if revealSkillAnswers {
				fmt.Printf("     answer: %s\n", q.answer)
			}
			n++
		}
	}
	print("focus — feeds today's drill", focus)
	print("redo — skills you marked missed", redo)
	print("cross-topic — a different skill each", cross)
	if !revealSkillAnswers {
		fmt.Println("  check: go run . -- --show     then: --missed=skill,skill  /  --got=skill")
	}
}
