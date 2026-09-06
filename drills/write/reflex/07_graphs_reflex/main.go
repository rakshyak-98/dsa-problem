// REFLEX DRILL 07 — Graphs (Grid BFS/DFS)
//
// RUN: go run -C drills/write/reflex/07_graphs_reflex .
//
// AFTER PASSING: graphs/medium/number_of_islands.js
package main

// TODO: REFLEX — count islands ('1' land, '0' water)
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	count := 0
	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	for r := range rows {
		for c := range cols {
			if grid[r][c] == '0' {
				continue
			}
			count++
			stack := [][]int{{r, c}}
			for len(stack) > 0 {
				current := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				cr, cc := current[0], current[1]
				for _, d := range directions {
					nr, nc := cr+d[0], cc+d[1]
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
						continue
					}
					if grid[nr][nc] == '0' {
						continue
					}
					grid[nr][nc] = '0'
					stack = append(stack, []int{nr, nc})
				}
			}
		}
	}
	return count
}

// TODO: REFLEX — flood fill from (sr, sc) with color
func floodFill(image [][]int, sr, sc, color int) [][]int {
	if len(image) == 0 {
		return image
	}
	originalColor := image[sr][sc]
	if originalColor == color {
		return image
	}
	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	stack := [][]int{{sr, sc}}
	image[sr][sc] = color
	rows, cols := len(image), len(image[0])

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		r, c := current[0], current[1]
		for _, d := range directions {
			nr, nc := r + d[0], c+d[1]

			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}
			image[nr][nc] = color
			stack = append(stack, []int{nr, nc})
		}
	}
	return image
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
