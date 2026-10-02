# Question Literacy Drills

Before writing code, practice **reading** a problem statement and filling in the ask template.

## The template (every problem)

```
Ask:        What is this asking in one sentence?
Input:      types, sorted?, duplicates?, negatives?
Output:     value / index / boolean / mutated array / count?
Forbidden:  extra space? modify input? recursion?
Edge:       empty, size 1, all same, max constraints
Pattern:    which template fits (after understanding, not before)
```

## How to use

1. Run `go run .` — today's ask is printed in the **Question Literacy** section.
2. Cover the hints and fill each bullet from the statement alone.
3. Only then open the matching reflex drill or primary problem.

## Small skills (same section, every day)

Under each ask sits a **SMALL SKILLS** block: short questions on the mechanics
the day's pattern is built from (index math, ranges, invariants, why a pointer
moves, what a variable means, complexity, and more). Four focus questions feed
today's drill; four cross-topic ones rotate through every other skill. Say the
answer *and the reason* first; `go run . -- --show` reveals them. The bank is
`skillBank` in `asks.go`; the curriculum is in `doc/write/STUDY_PLAN.md`.

## Rule

If you cannot write the **Ask** line without mentioning a data structure, you do not understand the problem yet.

## Weekday rotation

| Day | Theme | Primary after ask |
|-----|-------|-------------------|
| Mon | Array rotation | `arrays/easy/shuffle_the_array.js` |
| Tue | Two sum | `hashing/easy/two_sum.js` |
| Wed | Container water | `two_pointers/medium/container_with_most_water.js` |
| Thu | Search insert | `binary_search/easy/search_insertion_position.js` |
| Fri | Valid parens | `stacks/easy/valid_parentheses.js` |
| Sat | Min cost stairs | `dynamic_programming/easy/min_cost_climbing_staris.js` |
| Sun | Number of islands | `graphs/medium/number_of_islands.js` |

After the ask drill: run Core 5 → specialty reflex → solve the primary problem → check `go run . -- --levels`.
