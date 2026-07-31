# Drill 07 — Graphs

## numIslands
- **Pattern:** double loop; DFS/BFS each unvisited '1'
- **Mark visited** by mutating to '0' or visited set

## floodFill
- **Pattern:** DFS from (sr,sc) while same color
- **Bug:** forget bounds check

## shortestPathGrid
- **Pattern:** BFS from (0,0); first time reaching (n-1,m-1) is shortest
- **Return -1** if queue exhausts without reaching goal
