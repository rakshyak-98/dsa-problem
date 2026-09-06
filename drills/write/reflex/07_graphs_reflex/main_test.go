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

func TestCanFinish(t *testing.T) {
	assert(t, "canFinish no prereqs", canFinish(3, [][]int{}))
	assert(t, "canFinish chain", canFinish(3, [][]int{{1, 0}, {2, 1}}))
	assert(t, "canFinish simple cycle", !canFinish(2, [][]int{{1, 0}, {0, 1}}))
	assert(t, "canFinish self cycle", !canFinish(1, [][]int{{0, 0}}))
	assert(t, "canFinish diamond", canFinish(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	assert(t, "canFinish cycle in tail", !canFinish(4, [][]int{{1, 0}, {2, 1}, {1, 2}}))
}

// diamond adjacency list: 0->{1,2}, 1->{0,3}, 2->{0,3}, 3->{1,2}
var diamond = [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}

func TestDfs(t *testing.T) {
	assert(t, "dfs diamond", reflect.DeepEqual(dfs(diamond, 0), []int{0, 1, 3, 2}))
	assert(t, "dfs single", reflect.DeepEqual(dfs([][]int{{}}, 0), []int{0}))
	assert(t, "dfs disconnected", reflect.DeepEqual(dfs([][]int{{1}, {0}, {3}, {2}}, 0), []int{0, 1}))
	assert(t, "dfs line", reflect.DeepEqual(dfs([][]int{{1}, {2}, {3}, {}}, 0), []int{0, 1, 2, 3}))
	assert(t, "dfs star", reflect.DeepEqual(dfs([][]int{{1, 2, 3}, {0}, {0}, {0}}, 0), []int{0, 1, 2, 3}))
	assert(t, "dfs start not zero", reflect.DeepEqual(dfs([][]int{{1}, {2}, {0}}, 2), []int{2, 0, 1}))
}

func TestBfs(t *testing.T) {
	assert(t, "bfs diamond", reflect.DeepEqual(bfs(diamond, 0), []int{0, 1, 2, 3}))
	assert(t, "bfs single", reflect.DeepEqual(bfs([][]int{{}}, 0), []int{0}))
	assert(t, "bfs disconnected", reflect.DeepEqual(bfs([][]int{{1}, {0}, {3}, {2}}, 0), []int{0, 1}))
	assert(t, "bfs line", reflect.DeepEqual(bfs([][]int{{1}, {2}, {3}, {}}, 0), []int{0, 1, 2, 3}))
	assert(t, "bfs star", reflect.DeepEqual(bfs([][]int{{1, 2, 3}, {0}, {0}, {0}}, 0), []int{0, 1, 2, 3}))
	assert(t, "bfs start not zero", reflect.DeepEqual(bfs([][]int{{1}, {2}, {0}}, 2), []int{2, 0, 1}))
}

func TestBfsShortestPath(t *testing.T) {
	assert(t, "bfsShortestPath diamond", bfsShortestPath(diamond, 0, 3) == 2)
	assert(t, "bfsShortestPath same node", bfsShortestPath(diamond, 0, 0) == 0)
	assert(t, "bfsShortestPath unreachable", bfsShortestPath([][]int{{1}, {0}, {3}, {2}}, 0, 3) == -1)
	assert(t, "bfsShortestPath line", bfsShortestPath([][]int{{1}, {2}, {3}, {}}, 0, 3) == 3)
	assert(t, "bfsShortestPath through hub", bfsShortestPath([][]int{{1, 2, 3}, {0}, {0}, {0}}, 1, 2) == 2)
	assert(t, "bfsShortestPath directed cycle", bfsShortestPath([][]int{{1}, {2}, {0}}, 2, 1) == 2)
}

func TestTopoSort(t *testing.T) {
	assert(t, "topoSort diamond dag", reflect.DeepEqual(topoSort(4, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}), []int{0, 1, 2, 3}))
	assert(t, "topoSort chain", reflect.DeepEqual(topoSort(2, [][]int{{0, 1}}), []int{0, 1}))
	assert(t, "topoSort cycle", len(topoSort(3, [][]int{{0, 1}, {1, 2}, {2, 0}})) == 0)
	assert(t, "topoSort no edges", reflect.DeepEqual(topoSort(3, [][]int{}), []int{0, 1, 2}))
	assert(t, "topoSort single", reflect.DeepEqual(topoSort(1, [][]int{}), []int{0}))
	assert(t, "topoSort course schedule", reflect.DeepEqual(
		topoSort(6, [][]int{{5, 2}, {5, 0}, {4, 0}, {4, 1}, {2, 3}, {3, 1}}),
		[]int{4, 5, 0, 2, 3, 1}))
}

func TestDijkstra(t *testing.T) {
	g := [][]int{{0, 1, 4}, {0, 2, 1}, {2, 1, 2}, {1, 3, 1}, {2, 3, 5}}
	assert(t, "dijkstra weighted", reflect.DeepEqual(dijkstra(5, g, 0), []int{0, 3, 1, 4, -1}))
	assert(t, "dijkstra single", reflect.DeepEqual(dijkstra(1, [][]int{}, 0), []int{0}))
	assert(t, "dijkstra unreachable", reflect.DeepEqual(dijkstra(3, [][]int{{0, 1, 10}}, 0), []int{0, 10, -1}))
	assert(t, "dijkstra parallel edges", reflect.DeepEqual(dijkstra(2, [][]int{{0, 1, 5}, {0, 1, 3}}, 0), []int{0, 3}))
	assert(t, "dijkstra line", reflect.DeepEqual(dijkstra(4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}}, 0), []int{0, 1, 2, 3}))
	assert(t, "dijkstra from other source", reflect.DeepEqual(dijkstra(3, [][]int{{0, 1, 2}, {1, 2, 3}}, 0), []int{0, 2, 5}))
}

func TestAll(t *testing.T) {
	t.Run("numIslands", TestNumIslands)
	t.Run("floodFill", TestFloodFill)
	t.Run("shortestPathGrid", TestShortestPathGrid)
	t.Run("canFinish", TestCanFinish)
	t.Run("dfs", TestDfs)
	t.Run("bfs", TestBfs)
	t.Run("bfsShortestPath", TestBfsShortestPath)
	t.Run("topoSort", TestTopoSort)
	t.Run("dijkstra", TestDijkstra)
	fmt.Println("\nAll graph reflex drills passed.")
}
