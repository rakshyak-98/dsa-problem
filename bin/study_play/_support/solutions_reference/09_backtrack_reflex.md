# Drill 09 — Backtracking (bonus)

## subsets
- **Pattern:** for each index, branch include / exclude
- **Base:** path length == n or index == n

## permute
- **Pattern:** swap / choose unused; backtrack after recurse
- **Base:** path length == n

## combine
- **Pattern:** start from 1..n, pick next > last chosen
- **Prune** when remaining numbers can't fill k slots
