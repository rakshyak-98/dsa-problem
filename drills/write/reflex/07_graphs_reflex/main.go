// REFLEX DRILL 07 — Graphs (Grid BFS/DFS)
//
// RUN: go run -C drills/write/reflex/07_graphs_reflex .
//
// AFTER PASSING: graphs/medium/number_of_islands.js
package main

import (
	"fmt"
	"reflect"
)

// TODO: REFLEX — count islands ('1' land, '0' water)
func numIslands(grid [][]byte) int {
	panic("Implement from memory")
}

// TODO: REFLEX — flood fill from (sr, sc) with color
func floodFill(image [][]int, sr, sc, color int) [][]int {
	panic("Implement from memory")
}

// TODO: REFLEX — BFS shortest path length in unweighted grid (0=walkable, 1=wall)
// Return -1 if no path from top-left to bottom-right
func shortestPathGrid(grid [][]int) int {
	panic("Implement from memory")
}

func assert(name string, cond bool) {
	if !cond {
		panic(fmt.Sprintf("FAIL: %s", name))
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

func main() {
	grid1 := [][]byte{
		{'1', '1', '0'},
		{'0', '1', '0'},
		{'1', '0', '1'},
	}

	// numIslands — multi, all water/land, empty, single, diagonal, row
	assert("numIslands basic", numIslands(grid1) == 3)
	assert("numIslands all water", numIslands([][]byte{{'0'}}) == 0)
	assert("numIslands all land", numIslands([][]byte{{'1', '1'}, {'1', '1'}}) == 1)
	assert("numIslands empty", numIslands([][]byte{}) == 0)
	assert("numIslands single land", numIslands([][]byte{{'1'}}) == 1)
	assert("numIslands diagonal", numIslands([][]byte{{'1', '0'}, {'0', '1'}}) == 2)
	assert("numIslands row", numIslands([][]byte{{'1', '1', '1', '0', '1'}}) == 2)

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

	// floodFill — region, same color, single pixel, corner, isolated pixel
	assert("floodFill basic", reflect.DeepEqual(floodFill(cloneGrid(img), 1, 1, 2), want))
	sameColor := [][]int{{3}}
	assert("floodFill same color", reflect.DeepEqual(floodFill(cloneGrid(sameColor), 0, 0, 3), sameColor))
	singlePixel := [][]int{{0}}
	assert("floodFill single", reflect.DeepEqual(floodFill(cloneGrid(singlePixel), 0, 0, 9), [][]int{{9}}))
	corner := [][]int{{1, 0}, {0, 0}}
	assert("floodFill corner", reflect.DeepEqual(floodFill(cloneGrid(corner), 0, 0, 7), [][]int{{7, 0}, {0, 0}}))
	isolated := [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}
	assert("floodFill isolated", reflect.DeepEqual(floodFill(cloneGrid(isolated), 1, 1, 9), [][]int{{1, 0, 1}, {0, 9, 0}, {1, 0, 1}}))

	pathGrid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
	}

	// shortestPathGrid — detour, blocked, single, start/end blocked, straight, end wall
	assert("shortestPathGrid detour", shortestPathGrid(pathGrid) == 4)
	assert("shortestPathGrid blocked", shortestPathGrid([][]int{{0, 1}, {1, 0}}) == -1)
	assert("shortestPathGrid single", shortestPathGrid([][]int{{0}}) == 1)
	assert("shortestPathGrid start blocked", shortestPathGrid([][]int{{1}}) == -1)
	openPath := [][]int{{0, 0, 0, 0}}
	assert("shortestPathGrid straight", shortestPathGrid(openPath) == 3)
	endBlocked := [][]int{{0, 0}, {0, 1}}
	assert("shortestPathGrid end blocked", shortestPathGrid(endBlocked) == -1)
	open2x2 := [][]int{{0, 0}, {0, 0}}
	assert("shortestPathGrid open 2x2", shortestPathGrid(open2x2) == 2)

	fmt.Println("\nAll graph reflex drills passed.")
	fmt.Println("Primary: graphs/medium/number_of_islands.js")
}
