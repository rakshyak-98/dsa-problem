// REFLEX DRILL 07 — Graphs (Grid BFS/DFS)
//
// RUN: go run ./drills/07_graphs_reflex
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
	assert("numIslands", numIslands(grid1) == 3)
	assert("numIslands all water", numIslands([][]byte{{'0'}}) == 0)
	assert("numIslands all land", numIslands([][]byte{{'1', '1'}, {'1', '1'}}) == 1)

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
	assert("floodFill", reflect.DeepEqual(floodFill(cloneGrid(img), 1, 1, 2), want))

	pathGrid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
	}
	assert("shortestPathGrid", shortestPathGrid(pathGrid) == 4)
	assert("shortestPathGrid blocked", shortestPathGrid([][]int{{0, 1}, {1, 0}}) == -1)
	assert("shortestPathGrid single", shortestPathGrid([][]int{{0}}) == 1)

	fmt.Println("\nAll graph reflex drills passed.")
	fmt.Println("Primary: graphs/medium/number_of_islands.js")
}
