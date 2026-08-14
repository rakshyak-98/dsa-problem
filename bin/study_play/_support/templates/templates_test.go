package templates

import (
	"reflect"
	"testing"
)

func cloneBytes(grid [][]byte) [][]byte {
	out := make([][]byte, len(grid))
	for i, row := range grid {
		out[i] = append([]byte(nil), row...)
	}
	return out
}

func cloneInts(grid [][]int) [][]int {
	out := make([][]int, len(grid))
	for i, row := range grid {
		out[i] = append([]int(nil), row...)
	}
	return out
}

func TestArrays(t *testing.T) {
	if got := ReverseInPlaceTemplate([]int{1, 2, 3}); !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Fatalf("reverse: %v", got)
	}
	if got := ReverseInPlaceTemplate([]int{}); !reflect.DeepEqual(got, []int{}) {
		t.Fatalf("reverse empty: %v", got)
	}
	if IndexOfMaxTemplate([]int{3, 1, 4, 4}) != 2 {
		t.Fatal("indexOfMax")
	}
	if IndexOfMaxTemplate([]int{}) != -1 {
		t.Fatal("indexOfMax empty")
	}
	if ArraySumTemplate([]int{1, 2, 3}) != 6 {
		t.Fatal("arraySum")
	}
	if !reflect.DeepEqual(RotateRightTemplate([]int{1, 2, 3, 4, 5}, 2), []int{4, 5, 1, 2, 3}) {
		t.Fatal("rotateRight")
	}
	if !reflect.DeepEqual(RotateRightTemplate([]int{1, 2}, 5), []int{2, 1}) {
		t.Fatal("rotateRight k>len")
	}
	if !reflect.DeepEqual(RunningSumTemplate([]int{1, 2, 3}), []int{1, 3, 6}) {
		t.Fatal("runningSum")
	}
}

func TestHashing(t *testing.T) {
	if got := TwoSumTemplate([]int{2, 7, 11, 15}, 9); !reflect.DeepEqual(got, []int{0, 1}) {
		t.Fatalf("twoSum: %v", got)
	}
	if !ContainsDuplicateTemplate([]int{1, 2, 3, 1}) {
		t.Fatal("containsDuplicate true")
	}
	if ContainsDuplicateTemplate([]int{1, 2, 3}) {
		t.Fatal("containsDuplicate false")
	}
	freq := FrequencyMapTemplate([]string{"a", "b", "a"})
	if freq["a"] != 2 || freq["b"] != 1 {
		t.Fatal("frequencyMap")
	}
	if FirstUniqueCharTemplate("leetcode") != "l" {
		t.Fatal("firstUniqueChar")
	}
	if FirstUniqueCharTemplate("aabb") != "" {
		t.Fatal("firstUniqueChar none")
	}
	groups := GroupAnagramsTemplate([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	if len(groups) != 3 {
		t.Fatalf("groupAnagrams: %d", len(groups))
	}
}

func TestTwoPointers(t *testing.T) {
	dup := []int{1, 1, 2, 2, 3}
	if RemoveDuplicatesTemplate(dup) != 3 || !reflect.DeepEqual(dup[:3], []int{1, 2, 3}) {
		t.Fatal("removeDuplicates")
	}
	zeros := []int{0, 1, 0, 3, 12}
	MoveZeroesTemplate(zeros)
	if !reflect.DeepEqual(zeros, []int{1, 3, 12, 0, 0}) {
		t.Fatal("moveZeroes")
	}
	if MaxAreaTemplate([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) != 49 {
		t.Fatal("maxArea")
	}
	if !IsPalindromeTemplate("A man, a plan, a canal: Panama") {
		t.Fatal("isPalindrome true")
	}
	if IsPalindromeTemplate("race a car") {
		t.Fatal("isPalindrome false")
	}
	if MaxSumSubarrayKTemplate([]int{2, 1, 5, 1, 3, 2}, 3) != 9 {
		t.Fatal("maxSumSubarrayK")
	}
	if MaxSumSubarrayKTemplate([]int{1}, 2) != 0 {
		t.Fatal("maxSumSubarrayK invalid k")
	}
}

func TestBinarySearch(t *testing.T) {
	if BinarySearchTemplate([]int{-1, 0, 3, 5, 9, 12}, 9) != 4 {
		t.Fatal("binarySearch found")
	}
	if BinarySearchTemplate([]int{-1, 0, 3, 5, 9, 12}, 2) != -1 {
		t.Fatal("binarySearch missing")
	}
	if SearchInsertTemplate([]int{1, 3, 5, 6}, 2) != 1 {
		t.Fatal("searchInsert")
	}
	if FindMinRotatedTemplate([]int{4, 5, 6, 7, 0, 1, 2}) != 0 {
		t.Fatal("findMinRotated")
	}
	if !IsTargetPresentTemplate([]int{1, 2, 3}, 2) {
		t.Fatal("isTargetPresent")
	}
}

func TestTreesStacks(t *testing.T) {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}
	if !reflect.DeepEqual(InorderTraversalTemplate(tree), []int{2, 1, 4, 3}) {
		t.Fatal("inorder")
	}
	if !reflect.DeepEqual(InorderTraversalTemplate(nil), []int{}) {
		t.Fatal("inorder nil")
	}
	if !reflect.DeepEqual(PreorderTraversalTemplate(tree), []int{1, 2, 3, 4}) {
		t.Fatal("preorder")
	}
	if !reflect.DeepEqual(PostorderTraversalTemplate(tree), []int{2, 4, 3, 1}) {
		t.Fatal("postorder")
	}
	if !reflect.DeepEqual(LevelOrderTraversalTemplate(tree), []int{1, 2, 3, 4}) {
		t.Fatal("levelOrder")
	}
	if MaxDepthTemplate(tree) != 3 {
		t.Fatal("maxDepth")
	}
	if !IsValidParenthesesTemplate("()[]{}") || IsValidParenthesesTemplate("(]") {
		t.Fatal("validParens")
	}
	want := []int{1, 1, 4, 2, 1, 1, 0, 0}
	if !reflect.DeepEqual(DailyTemperaturesTemplate([]int{73, 74, 75, 71, 69, 72, 76, 73}), want) {
		t.Fatal("dailyTemperatures")
	}
}

func TestDP(t *testing.T) {
	if FibTemplate(10) != 55 || FibTemplate(0) != 0 {
		t.Fatal("fib")
	}
	if MinCostClimbingStairsTemplate([]int{10, 15, 20}) != 15 {
		t.Fatal("minCost")
	}
	if RobTemplate([]int{2, 7, 9, 3, 1}) != 12 {
		t.Fatal("rob")
	}
	if ClimbStairsTemplate(5) != 8 {
		t.Fatal("climbStairs")
	}
}

func TestGraphs(t *testing.T) {
	grid := [][]byte{{'1', '1', '0'}, {'0', '1', '0'}, {'1', '0', '1'}}
	if NumIslandsTemplate(cloneBytes(grid)) != 3 {
		t.Fatal("numIslands")
	}
	img := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	got := FloodFillTemplate(cloneInts(img), 1, 1, 2)
	if got[0][0] != 2 || got[2][2] != 1 {
		t.Fatal("floodFill")
	}
	same := FloodFillTemplate(cloneInts([][]int{{1}}), 0, 0, 1)
	if same[0][0] != 1 {
		t.Fatal("floodFill same color")
	}
	path := [][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}}
	if ShortestPathGridTemplate(path) != 4 {
		t.Fatal("shortestPath")
	}
	if ShortestPathGridTemplate([][]int{{0, 1}, {1, 0}}) != -1 {
		t.Fatal("shortestPath blocked")
	}
}

func TestVariants(t *testing.T) {
	if !reflect.DeepEqual(TwoSumSortedTemplate([]int{2, 7, 11, 15}, 9), []int{0, 1}) {
		t.Fatal("twoSumSorted")
	}
	if MaxSubarraySumTemplate([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) != 6 {
		t.Fatal("maxSubarraySum")
	}
	if LengthOfLongestSubstringTemplate("abcabcbb") != 3 {
		t.Fatal("longestSubstring")
	}
	if !reflect.DeepEqual(ProductExceptSelfTemplate([]int{1, 2, 3, 4}), []int{24, 12, 8, 6}) {
		t.Fatal("productExceptSelf")
	}
}

func TestHeaps(t *testing.T) {
	if KthLargestTemplate([]int{3, 2, 1, 5, 6, 4}, 2) != 5 {
		t.Fatal("kthLargest")
	}
	if LastStoneWeightTemplate([]int{2, 7, 4, 1, 8, 1}) != 1 {
		t.Fatal("lastStoneWeight")
	}
	got := MergeKSortedTemplate([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}})
	if len(got) != 8 || got[0] != 1 || got[len(got)-1] != 6 {
		t.Fatalf("mergeKSorted: %v", got)
	}
	if len(MergeKSortedTemplate(nil)) != 0 {
		t.Fatal("mergeKSorted empty")
	}
}

func TestBacktrack(t *testing.T) {
	if len(SubsetsTemplate([]int{1, 2, 3})) != 8 {
		t.Fatal("subsets")
	}
	if len(PermuteTemplate([]int{1, 2, 3})) != 6 {
		t.Fatal("permute")
	}
	if len(CombineTemplate(4, 2)) != 6 {
		t.Fatal("combine")
	}
}

func TestCore5Templates(t *testing.T) {
	if got := TwoSumTemplate([]int{-1, -2, -3, -4, -5}, -8); !reflect.DeepEqual(got, []int{2, 4}) {
		t.Fatalf("core5 twoSum negatives: %v", got)
	}
	dup := []int{1, 1, 2}
	if RemoveDuplicatesTemplate(dup) != 2 {
		t.Fatal("core5 removeDuplicates")
	}
	if BinarySearchTemplate([]int{}, 1) != -1 {
		t.Fatal("core5 binarySearch empty")
	}
	if len(FrequencyMapTemplate([]string{})) != 0 {
		t.Fatal("core5 frequencyMap empty")
	}
}
