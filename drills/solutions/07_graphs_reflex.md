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

## canFinish
- **Trigger:** prerequisites, ordering, "is there a cycle" on a directed graph
- **Pattern:** Kahn — in-degree array + queue of zero-in-degree nodes
- **Cycle test:** if fewer than numCourses nodes come off the queue, one remains
- **Bug:** building edges backwards — `[course, need]` means an edge need → course

## dfs / bfs (adjacency list)
- **dfs:** recursive preorder — mark on entry, append, recurse unseen neighbours.
  Marking *after* the recursion revisits nodes on a cycle.
- **bfs:** queue, mark **on enqueue** — mark on dequeue and a node enters twice.

## bfsShortestPath
- **Trigger:** "fewest steps / shortest path" on an unweighted graph
- **Pattern:** level-by-level BFS; the count of levels crossed is the edge count
- **Never DFS for this** — DFS finds *a* path, not the shortest

## topoSort
- **Trigger:** directed graph, valid order, or "is there a cycle"
- **Pattern:** Kahn — queue zero-in-degree nodes (smallest index on ties), pop, decrement targets
- **Cycle:** fewer than n emitted → nil

## dijkstra
- **Trigger:** cheapest weighted path, non-negative edges
- **Pattern:** min-heap of `(node, dist)`; pop closest, relax edges, skip stale `(v,d)` where `d > dist[v]`
- **Negative edges** need Bellman–Ford instead
