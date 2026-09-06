// SOLUTION — Reflex 07 Graphs (peek after honest attempt)
package main

import (
	"container/heap"
	"fmt"
	"reflect"
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	count := 0

	var sink func(r, c int)
	sink = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0'
		sink(r+1, c)
		sink(r-1, c)
		sink(r, c+1)
		sink(r, c-1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				sink(r, c)
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

func canFinish(numCourses int, prerequisites [][]int) bool {
	// Kahn: build an adjacency list plus in-degrees, then repeatedly take a
	// node with no unmet prerequisite. A leftover node means a cycle.
	next := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		course, need := p[0], p[1]
		next[need] = append(next[need], course)
		indeg[course]++
	}
	queue := []int{}
	for i, d := range indeg {
		if d == 0 {
			queue = append(queue, i)
		}
	}
	done := 0
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		done++
		for _, m := range next[n] {
			indeg[m]--
			if indeg[m] == 0 {
				queue = append(queue, m)
			}
		}
	}
	return done == numCourses
}

// dfs — recursive preorder over an adjacency list. Mark on entry, append,
// recurse into unseen neighbours in list order.
func dfs(graph [][]int, start int) []int {
	seen := make([]bool, len(graph))
	var order []int
	var visit func(u int)
	visit = func(u int) {
		seen[u] = true
		order = append(order, u)
		for _, v := range graph[u] {
			if !seen[v] {
				visit(v)
			}
		}
	}
	visit(start)
	return order
}

// bfs — queue, mark on enqueue (never on dequeue, or a node enters twice).
func bfs(graph [][]int, start int) []int {
	seen := make([]bool, len(graph))
	seen[start] = true
	q := []int{start}
	var order []int
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		order = append(order, u)
		for _, v := range graph[u] {
			if !seen[v] {
				seen[v] = true
				q = append(q, v)
			}
		}
	}
	return order
}

// bfsShortestPath — level-by-level BFS; the number of levels crossed is the
// edge count. Unweighted shortest path is BFS, never DFS.
func bfsShortestPath(graph [][]int, src, dst int) int {
	if src == dst {
		return 0
	}
	seen := make([]bool, len(graph))
	seen[src] = true
	q := []int{src}
	dist := 0
	for len(q) > 0 {
		dist++
		var next []int
		for _, u := range q {
			for _, v := range graph[u] {
				if v == dst {
					return dist
				}
				if !seen[v] {
					seen[v] = true
					next = append(next, v)
				}
			}
		}
		q = next
	}
	return -1
}

// topoSort — Kahn. Repeatedly take a zero-in-degree node (smallest index on
// ties), remove its out-edges. Fewer than n emitted means a cycle blocked it.
func topoSort(n int, edges [][]int) []int {
	adj := make([][]int, n)
	indeg := make([]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		indeg[e[1]]++
	}
	order := make([]int, 0, n)
	used := make([]bool, n)
	for len(order) < n {
		pick := -1
		for v := 0; v < n; v++ {
			if !used[v] && indeg[v] == 0 {
				pick = v
				break
			}
		}
		if pick == -1 {
			return nil
		}
		used[pick] = true
		order = append(order, pick)
		for _, v := range adj[pick] {
			indeg[v]--
		}
	}
	return order
}

type distNode struct{ v, d int }

type distHeap []distNode

func (h distHeap) Len() int            { return len(h) }
func (h distHeap) Less(i, j int) bool  { return h[i].d < h[j].d }
func (h distHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *distHeap) Push(x interface{}) { *h = append(*h, x.(distNode)) }
func (h *distHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

// dijkstra — greedy shortest paths, non-negative weights. Pop the closest
// frontier node, relax its edges, skip stale heap entries (d > dist[v]).
func dijkstra(n int, edges [][]int, src int) []int {
	adj := make([][][2]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], [2]int{e[1], e[2]})
	}
	const inf = 1 << 60
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[src] = 0
	pq := &distHeap{{src, 0}}
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(distNode)
		if cur.d > dist[cur.v] {
			continue
		}
		for _, nb := range adj[cur.v] {
			if nd := cur.d + nb[1]; nd < dist[nb[0]] {
				dist[nb[0]] = nd
				heap.Push(pq, distNode{nb[0], nd})
			}
		}
	}
	for i := range dist {
		if dist[i] == inf {
			dist[i] = -1
		}
	}
	return dist
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

	// canFinish — no edges, chain, self cycle, two-node cycle, diamond
	assert("canFinish no prereqs", canFinish(3, [][]int{}))
	assert("canFinish chain", canFinish(3, [][]int{{1, 0}, {2, 1}}))
	assert("canFinish simple cycle", !canFinish(2, [][]int{{1, 0}, {0, 1}}))
	assert("canFinish self cycle", !canFinish(1, [][]int{{0, 0}}))
	assert("canFinish diamond", canFinish(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	assert("canFinish cycle in tail", !canFinish(4, [][]int{{1, 0}, {2, 1}, {1, 2}}))

	// dfs / bfs — diamond adjacency list, disconnected, start not zero
	diamond := [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}
	assert("dfs diamond", reflect.DeepEqual(dfs(diamond, 0), []int{0, 1, 3, 2}))
	assert("dfs disconnected", reflect.DeepEqual(dfs([][]int{{1}, {0}, {3}, {2}}, 0), []int{0, 1}))
	assert("bfs diamond", reflect.DeepEqual(bfs(diamond, 0), []int{0, 1, 2, 3}))
	assert("bfs star", reflect.DeepEqual(bfs([][]int{{1, 2, 3}, {0}, {0}, {0}}, 0), []int{0, 1, 2, 3}))

	// bfsShortestPath — reachable, same node, unreachable
	assert("bfsShortestPath diamond", bfsShortestPath(diamond, 0, 3) == 2)
	assert("bfsShortestPath same", bfsShortestPath(diamond, 0, 0) == 0)
	assert("bfsShortestPath unreachable", bfsShortestPath([][]int{{1}, {0}, {3}, {2}}, 0, 3) == -1)

	// topoSort — DAG, cycle, course-schedule ordering
	assert("topoSort dag", reflect.DeepEqual(topoSort(4, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}), []int{0, 1, 2, 3}))
	assert("topoSort cycle", len(topoSort(3, [][]int{{0, 1}, {1, 2}, {2, 0}})) == 0)
	assert("topoSort course schedule", reflect.DeepEqual(
		topoSort(6, [][]int{{5, 2}, {5, 0}, {4, 0}, {4, 1}, {2, 3}, {3, 1}}),
		[]int{4, 5, 0, 2, 3, 1}))

	// dijkstra — weighted, unreachable, parallel edges
	dg := [][]int{{0, 1, 4}, {0, 2, 1}, {2, 1, 2}, {1, 3, 1}, {2, 3, 5}}
	assert("dijkstra weighted", reflect.DeepEqual(dijkstra(5, dg, 0), []int{0, 3, 1, 4, -1}))
	assert("dijkstra unreachable", reflect.DeepEqual(dijkstra(3, [][]int{{0, 1, 10}}, 0), []int{0, 10, -1}))
	assert("dijkstra parallel edges", reflect.DeepEqual(dijkstra(2, [][]int{{0, 1, 5}, {0, 1, 3}}, 0), []int{0, 3}))

	fmt.Println("\nAll graph reflex drills passed.")
	fmt.Println("Primary: graphs/medium/number_of_islands.js")
}
