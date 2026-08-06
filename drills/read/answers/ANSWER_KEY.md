# Answer key — peek only after `go run` fails

Do **not** memorize these. Re-trace the snippet after reading a line.

---

## 00_core_read

| Var | Answer |
|-----|--------|
| `aPattern` | `fixed sliding window` |
| `aTrace` | `9` (windows: 2+1+5=8, 1+5+1=7, 5+1+3=9, 1+3+2=6) |
| `aTime` | `O(n)` |
| `bPattern` | `hash complement` / `map two sum` |
| `bTrace` | `0,1` |
| `bMutates` | `no` |
| `cPattern` | `two pointers opposite ends` |
| `cTrace` | `true` |
| `cAsk` | e.g. check whether the string is a palindrome |

---

## 01_scan_structure

| Var | Answer |
|-----|--------|
| `s1Returns` | `index` |
| `s1Empty` | `-1` |
| `s1Loops` | `1` |
| `s2Mutates` | `no` |
| `s2Builds` | `prefix` |
| `s2Special` | `empty` |
| `s3Mutates` | `yes` |
| `s3Outer` | `nested loops` |
| `s3Helper` | `dfs` |

---

## 02_trace_execution

| Var | Answer |
|-----|--------|
| `t1Len` | `3` |
| `t1Prefix` | `1,2,3` |
| `t2Best` | `49` |
| `t3Mids` | `2,3` |
| `t3Ret` | `3` |

---

## 03_name_the_pattern

| Var | Answer |
|-----|--------|
| `pAlpha` | `variable sliding window` |
| `pBeta` | `monotonic stack` |
| `pGamma` | `binary search on rotated array` |
| `pDelta` | `1D DP` |
| `pEpsilon` | `two pointers` / `in-place reverse` |

---

## 04_find_the_bug

| Var | Answer |
|-----|--------|
| `bug1Kind` | `off-by-one` (`lo < hi` → `lo <= hi`) |
| `bug1Fix` | must mention `<=` |
| `bug2Kind` | `wrong-move` |
| `bug2Move` | `left` |
| `bug3Kind` | `missing-visit` |
| `bug3Missing` | `left` |

---

## 05_complexity_glance

| Var | Answer |
|-----|--------|
| `c1Time` | `O(n)` (amortized) |
| `c1Space` | `O(1)` |
| `c2Time` | `O(n^3)` |
| `c3Time` | `O(n)` |
| `c3Space` | `O(n)` |
| `c4Time` | `O(2^n)` |

---

## 06_reconstruct_ask

| Var | Idea |
|-----|------|
| `askA` | true if any duplicate value appears within distance k |
| `askB` | remove previous character for each `*` (like backspace) |
| `askC` | max sum of values with no two adjacent (house robber) |
| `askD` | whether two strings are anagrams |

---

## 07_compare_variants

| Var | Answer |
|-----|--------|
| `pair1Ask` | indices of two numbers that sum to target |
| `pair1Trade` | `n^2 vs n` |
| `pair2Ask` | ways to climb n stairs |
| `pair2Pick` | `v2b` |
| `pair3Mutates` | `yes` |
| `pair3Space` | `v3a` |

---

## 08_math_concepts

| Var | Answer |
|-----|--------|
| `m1GCD` | `6` |
| `m2ModPow` | `81` |
| `m3nCr` | `56` |
| `m4Sum` | `210` |
| `m5Fib` | `13` |
| `m6Master` | `O(n log n)` |
| `m7Recurrence` | `O(log n)` |
| `m8Bits` | `3` |
| `m9Prime` | `no` |
| `m10Height` | `3` |
