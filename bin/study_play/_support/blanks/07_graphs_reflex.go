//go:build ignore

// REFLEX DRILL 07 — Graphs (Grid BFS/DFS)
//
// RUN: go run -C drills/write/reflex/07_graphs_reflex .
//
// AFTER PASSING: graphs/medium/number_of_islands.js
package main

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

// TODO: REFLEX — can every course be finished (topological sort / cycle check)
func canFinish(numCourses int, prerequisites [][]int) bool {
	panic("Implement from memory")
}

func main() {}
