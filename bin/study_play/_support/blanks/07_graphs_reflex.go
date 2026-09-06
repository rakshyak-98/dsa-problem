//go:build ignore

// REFLEX DRILL 07 — Graphs (grid + adjacency-list algorithms)
//
// RUN: go run -C drills/write/reflex/07_graphs_reflex .
//
// The last five drills work on an adjacency list: graph[u] holds u's
// neighbours, in order. topoSort / dijkstra take n + an edge list.
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

// TODO: REFLEX — DFS visit order from start (recursive preorder, neighbours in list order)
func dfs(graph [][]int, start int) []int {
	panic("Implement from memory")
}

// TODO: REFLEX — BFS visit order from start (level by level, neighbours in list order)
func bfs(graph [][]int, start int) []int {
	panic("Implement from memory")
}

// TODO: REFLEX — fewest edges from src to dst; 0 if equal, -1 if unreachable
func bfsShortestPath(graph [][]int, src, dst int) int {
	panic("Implement from memory")
}

// TODO: REFLEX — Kahn topological order of a directed graph (edges u->v);
// smallest node index first on ties; nil if the graph has a cycle
func topoSort(n int, edges [][]int) []int {
	panic("Implement from memory")
}

// TODO: REFLEX — Dijkstra shortest distances from src over directed edges
// [u, v, w] with w >= 0; dist[src] = 0, unreachable = -1
func dijkstra(n int, edges [][]int, src int) []int {
	panic("Implement from memory")
}

func main() {}
