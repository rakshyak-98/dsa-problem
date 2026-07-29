# Reading Patterns — The 6-Pass Method

Use this on every snippet in the drills, and on editorials / your own past solutions.

Do **not** start at line 1 and hope. Scan with purpose.

---

## The 6-pass read

| Pass | Name | Time | Question you answer |
|------|------|------|---------------------|
| 1 | **Signature** | 15s | What goes in? What comes out? Mutates input? |
| 2 | **Skeleton** | 30s | Where are the loops / recursion / early returns? |
| 3 | **State** | 45s | Which variables change? What do they mean? |
| 4 | **Trace** | 2–3m | One concrete input → step values → final return |
| 5 | **Pattern** | 30s | Which known template is this? (map, L/R, window, BS, DFS, DP…) |
| 6 | **Ask + bound** | 45s | One-sentence problem + O(time) / O(space) |

If Pass 4 fails, go back to Pass 3 — you misnamed state.

---

## Pass details

### 1. Signature

Write (mentally or on paper):

```
in:  ...
out: ...
side effects: yes/no
```

Watch for: returns index vs value, bool vs count, in-place vs new slice.

### 2. Skeleton

Mark the control flow only:

- `for` / `for range` / `while`-style
- nested loop vs single pass
- recursion base cases
- `return` early exits

Ignore arithmetic until the skeleton is clear.

### 3. State

Name every binder that survives across iterations:

| Name in code | Role (your words) |
|--------------|-------------------|
| `left` | window start |
| `best` | answer so far |
| `seen` | values already visited |

Bad: “`i` is a counter.”  
Good: “`write` is the next free slot for kept elements.”

### 4. Trace

Pick the sample from the drill (or invent a 4–6 element case).

Table:

```
step | line/event        | key vars        | notes
0    | start             | left=0 best=0   |
1    | right=0 include x | ...             |
```

Predict the return **before** looking at any answer key.

### 5. Pattern

Match shape → name:

| Shape you see | Pattern name |
|---------------|--------------|
| `map` + `target - x` | hash complement |
| `left`/`right` moving inward | opposite two pointers |
| `write` advanced only when keep | read/write pointers |
| expand `right`, shrink `left` while invalid | variable sliding window |
| first window of `k`, then `+in -out` | fixed sliding window |
| `lo`/`hi`/`mid`, half discarded | binary search |
| stack of indices, pop while greater | monotonic stack |
| `dp[i]` from `dp[i-1]`… | 1D DP |
| queue + visited on grid | BFS |
| recurse 4 dirs / children | DFS |

### 6. Ask + bound

- Ask: one sentence a junior could understand (no jargon).
- Time: count nested loops / recursion branching carefully.
- Space: maps, stacks, recursion depth, extra arrays.

---

## Common reading traps

| Trap | Fix |
|------|-----|
| Reading names instead of roles | Rename state in Pass 3 (`i` → “read cursor”) |
| Skipping the sample | Always Pass 4 on at least one input |
| Confusing subarray vs subsequence from the code | Contiguous index range ⇒ subarray |
| Assuming sorted | Check whether code sorts or requires sorted input |
| Complexity from “looks nested” | Confirm inner loop actually runs O(n) times total (two pointers often O(n)) |

---

## Ownership criteria (reading)

You **own** a snippet when you can, without the answer key:

- [ ] State signature + mutation in 15s
- [ ] Name the pattern in one breath
- [ ] Trace one sample correctly
- [ ] Restate the ask in one sentence
- [ ] Give tight time/space bounds

Drill until the 6 passes feel automatic. Then speed will follow.
