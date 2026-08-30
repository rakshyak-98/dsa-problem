// REFLEX DRILL 07 — Graphs (Grid BFS/DFS)
//
// RUN: go run -C drills/write/reflex/07_graphs_reflex .
//
// AFTER PASSING: graphs/medium/number_of_islands.js
package main

// TODO: REFLEX — count islands ('1' land, '0' water)
func numIslands(grid [][]byte) int {
	count := 0
	if len(grid) == 0 {
		return count
	}
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	for r := range rows {
		for c := range cols {
			if grid[r][c] == '1' {
				count++
				grid[r][c] = '0' // mark visited
				stack := [][]int{{r, c}}
				for len(stack) > 0 {
					current := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					r := current[0]
					c := current[1]
					for _, d := range directions {
						nr := r + d[0]
						nc := c + d[1]
						if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] == '0' {
							continue
						}
						grid[nr][nc] = '0'
						stack = append(stack, []int{nr, nc})
					}
				}
			}
		}
	}
	return count
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

func main() {}
