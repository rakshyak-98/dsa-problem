// SOLUTION — Reflex 07 Graphs (peek after honest attempt)
package main

import (
	"fmt"
	"reflect"
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	count := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}
	return count
}

func floodFill(image [][]int, sr, sc, color int) [][]int {
	start := image[sr][sc]
	if start == color {
		return image
	}
	rows, cols := len(image), len(image[0])
	stack := [][2]int{{sr, sc}}
	for len(stack) > 0 {
		cell := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		r, c := cell[0], cell[1]
		if r < 0 || c < 0 || r >= rows || c >= cols || image[r][c] != start {
			continue
		}
		image[r][c] = color
		stack = append(stack, [2]int{r + 1, c}, [2]int{r - 1, c}, [2]int{r, c + 1}, [2]int{r, c - 1})
	}
	return image
}

func shortestPathGrid(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	if grid[0][0] == 1 || grid[rows-1][cols-1] == 1 {
		return -1
	}
	if rows == 1 && cols == 1 {
		return 1
	}
	type cell struct{ r, c, dist int }
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := []cell{{0, 0, 0}}
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	visited[0][0] = true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.r == rows-1 && cur.c == cols-1 {
			return cur.dist
		}
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr < 0 || nc < 0 || nr >= rows || nc >= cols {
				continue
			}
			if grid[nr][nc] == 1 || visited[nr][nc] {
				continue
			}
			visited[nr][nc] = true
			queue = append(queue, cell{nr, nc, cur.dist + 1})
		}
	}
	return -1
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
	assert("numIslands empty", numIslands([][]byte{}) == 0)
	assert("numIslands single land", numIslands([][]byte{{'1'}}) == 1)

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
	sameColor := [][]int{{3}}
	assert("floodFill same color", reflect.DeepEqual(floodFill(cloneGrid(sameColor), 0, 0, 3), sameColor))
	singlePixel := [][]int{{0}}
	assert("floodFill single", reflect.DeepEqual(floodFill(cloneGrid(singlePixel), 0, 0, 9), [][]int{{9}}))

	pathGrid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
	}
	assert("shortestPathGrid", shortestPathGrid(pathGrid) == 4)
	assert("shortestPathGrid blocked", shortestPathGrid([][]int{{0, 1}, {1, 0}}) == -1)
	assert("shortestPathGrid single", shortestPathGrid([][]int{{0}}) == 1)
	assert("shortestPathGrid start blocked", shortestPathGrid([][]int{{1}}) == -1)
	openPath := [][]int{{0, 0, 0, 0}}
	assert("shortestPathGrid straight", shortestPathGrid(openPath) == 3)

	fmt.Println("\nAll graph reflex drills passed.")
}
