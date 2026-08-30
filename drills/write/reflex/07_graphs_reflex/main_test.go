package main

import (
	"fmt"
	"reflect"
	"testing"
)

func assert(t *testing.T, name string, cond bool) {
	t.Helper()
	if !cond {
		fmt.Printf("FAIL: %s\n", name)
		t.Fatalf("FAIL: %s", name)
	}
	fmt.Printf("PASS: %s\n", name)
}

func cloneGrid(src [][]int) [][]int {
	out := make([][]int, len(src))
	for i := range src {
		out[i] = append([]int(nil), src[i]...)
	}
	return out
}

func TestNumIslands(t *testing.T) {
	grid1 := [][]byte{
		{'1', '1', '0'},
		{'0', '1', '0'},
		{'1', '0', '1'},
	}
	assert(t, "numIslands basic", numIslands(grid1) == 3)
	assert(t, "numIslands all water", numIslands([][]byte{{'0'}}) == 0)
	assert(t, "numIslands all land", numIslands([][]byte{{'1', '1'}, {'1', '1'}}) == 1)
	assert(t, "numIslands empty", numIslands([][]byte{}) == 0)
	assert(t, "numIslands single land", numIslands([][]byte{{'1'}}) == 1)
	assert(t, "numIslands diagonal", numIslands([][]byte{{'1', '0'}, {'0', '1'}}) == 2)
	assert(t, "numIslands row", numIslands([][]byte{{'1', '1', '1', '0', '1'}}) == 2)
}

func TestFloodFill(t *testing.T) {
	img := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	want := [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1},
	}
	assert(t, "floodFill basic", reflect.DeepEqual(floodFill(cloneGrid(img), 1, 1, 2), want))
	sameColor := [][]int{{3}}
	assert(t, "floodFill same color", reflect.DeepEqual(floodFill(cloneGrid(sameColor), 0, 0, 3), sameColor))
	singlePixel := [][]int{{0}}
	assert(t, "floodFill single", reflect.DeepEqual(floodFill(cloneGrid(singlePixel), 0, 0, 9), [][]int{{9}}))
	corner := [][]int{{1, 0}, {0, 0}}
	assert(t, "floodFill corner", reflect.DeepEqual(floodFill(cloneGrid(corner), 0, 0, 7), [][]int{{7, 0}, {0, 0}}))
	isolated := [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}
	assert(t, "floodFill isolated", reflect.DeepEqual(floodFill(cloneGrid(isolated), 1, 1, 9), [][]int{{1, 0, 1}, {0, 9, 0}, {1, 0, 1}}))
}

func TestShortestPathGrid(t *testing.T) {
	pathGrid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
	}
	assert(t, "shortestPathGrid detour", shortestPathGrid(pathGrid) == 4)
	assert(t, "shortestPathGrid blocked", shortestPathGrid([][]int{{0, 1}, {1, 0}}) == -1)
	assert(t, "shortestPathGrid single", shortestPathGrid([][]int{{0}}) == 1)
	assert(t, "shortestPathGrid start blocked", shortestPathGrid([][]int{{1}}) == -1)
	openPath := [][]int{{0, 0, 0, 0}}
	assert(t, "shortestPathGrid straight", shortestPathGrid(openPath) == 3)
	endBlocked := [][]int{{0, 0}, {0, 1}}
	assert(t, "shortestPathGrid end blocked", shortestPathGrid(endBlocked) == -1)
	open2x2 := [][]int{{0, 0}, {0, 0}}
	assert(t, "shortestPathGrid open 2x2", shortestPathGrid(open2x2) == 2)
}

func TestAll(t *testing.T) {
	t.Run("numIslands", TestNumIslands)
	t.Run("floodFill", TestFloodFill)
	t.Run("shortestPathGrid", TestShortestPathGrid)
	fmt.Println("\nAll graph reflex drills passed.")
}
