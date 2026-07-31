# DSA Math Concepts Guide

This file collects the **mathematical concepts, identities, and equations** that underpin data structures and algorithms. It is **not** tied to any single problem set — use it whenever you need to analyze complexity, count outcomes, reason about correctness, or choose the right numeric tool.

Sources distilled from standard references: *Introduction to Algorithms* (CLRS), *Algorithm Design* (Kleinberg & Tardos), *The Algorithm Design Manual* (Skiena), *Algorithms* (Sedgewick & Wayne), *Introduction to the Design and Analysis of Algorithms* (Levitin), and *The Art of Computer Programming* (Knuth).

---

## 01 — Asymptotic Complexity & Growth Rates

- **Core idea:** Describe how resource usage scales with input size `n`, ignoring constants and lower-order terms.
- **When this appears:** Every algorithm analysis — loops, recursion, data-structure operations, lower-bound arguments.
- **Key concepts:**

### Asymptotic notation

| Notation | Meaning | Intuition |
|----------|---------|-----------|
| **O(f(n))** | Upper bound (worst-case growth) | “At most ~f(n)” |
| **Ω(f(n))** | Lower bound | “At least ~f(n)” |
| **Θ(f(n))** | Tight bound | “Exactly ~f(n)” (both O and Ω) |
| **o(f(n))** | Strict upper bound | Grows strictly slower than f |
| **ω(f(n))** | Strict lower bound | Grows strictly faster than f |

**Formal (CLRS style):**  
`g(n) = O(f(n))` iff ∃ constants `c > 0`, `n₀` such that `0 ≤ g(n) ≤ c·f(n)` for all `n ≥ n₀`.

### Growth-rate hierarchy (increasing)

```
O(1) < O(log n) < O(√n) < O(n) < O(n log n) < O(n²) < O(n³) < O(2ⁿ) < O(n!) < O(nⁿ)
```

**Useful limits:**
- `logₐ n = (log_b n) / (log_b a)`  (change of base)
- `log(n!) = Θ(n log n)`  (Stirling: `n! ≈ √(2πn)(n/e)ⁿ`)
- `2^(log₂ n) = n`
- Polynomial `nᵏ` beats any `cⁿ` with `c > 1` for large enough `n`

### Common loop complexities

| Code shape | Time |
|------------|------|
| Single loop `i = 0..n` | O(n) |
| Nested `i,j` both `0..n` | O(n²) |
| Halving `n /= 2` each step | O(log n) |
| `i *= 2` until `n` | O(log n) |
| `for i=1; i<n; i*=2` + inner `j=0..i` | O(n) (geometric sum) |
| `for i=0; i<n; i++` + inner `j=0..i` | O(n²) |
| Three nested `0..n` | O(n³) |
| Divide array in half, one recursive call | O(n) total if O(1) work per level |
| Divide in half, two recursive calls + O(n) merge | O(n log n) |

### Amortized analysis

**Aggregate:** Total cost `T(n)` over `n` operations ⇒ amortized `T(n)/n`.

**Accounting method:** Assign credits to cheap ops to pay for expensive ones (e.g. dynamic array doubling: O(1) amortized insert).

**Potential method:** Define potential `Φ` on data structure; amortized cost = actual cost + ΔΦ.

**Classic results:**
- Dynamic array push: **O(1) amortized**
- Disjoint-set union–find with path compression + rank: **O(α(n))** per op (α = inverse Ackermann, effectively constant)
- Splay tree: **O(log n) amortized** per op

### Comparison-model lower bounds

- Sorting by comparisons: **Ω(n log n)** (decision tree has ≥ `n!` leaves ⇒ height ≥ log₂(n!))
- Searching sorted array: **Ω(log n)** comparisons
- Finding min/max: **Ω(n)** comparisons (can do both in ≤ `3n/2 - 2`)

---

## 02 — Summations, Series & Products

- **Core idea:** Closed forms turn nested loops and recursion trees into closed complexity bounds.
- **When this appears:** Analyzing loops, recursion trees, expected values, counting operations.

### Essential sums

| Sum | Closed form | Notes |
|-----|-------------|-------|
| `Σᵢ₌₁ⁿ 1` | `n` | |
| `Σᵢ₌₀ⁿ⁻¹ i` | `n(n-1)/2` | Arithmetic |
| `Σᵢ₌₁ⁿ i` | `n(n+1)/2` | |
| `Σᵢ₌₁ⁿ i²` | `n(n+1)(2n+1)/6` | |
| `Σᵢ₌₁ⁿ i³` | `[n(n+1)/2]²` | |
| `Σᵢ₌₀ⁿ rⁱ` | `(rⁿ⁺¹ - 1)/(r - 1)` | Geometric; `r≠1` |
| `Σᵢ₌₀^∞ rⁱ` | `1/(1-r)` | `\|r\| < 1` |
| `Σᵢ₌₁ⁿ 1/i` | `H(n) ≈ ln n + γ` | Harmonic; γ ≈ 0.577 |
| `Σᵢ₌₁ⁿ log i` | `Θ(n log n)` | |
| `Σᵢ₌₀^(log n) 2ⁱ` | `2^(log n + 1) - 1 = 2n - 1` | Geometric in tree height |

### Telescoping

`Σᵢ (aᵢ₊₁ - aᵢ) = aₙ - a₀`

Useful for simplifying recurrence expansions and some probability proofs.

### Products

- `Πᵢ₌₁ⁿ i = n!`
- `Πᵢ₌₁ⁿ c = cⁿ`
- Number of leaves in full binary tree of height `h`: `2ʰ`
- Nodes in perfect binary tree of height `h`: `2^(h+1) - 1`

---

## 03 — Logarithms & Exponents

- **Core idea:** Logarithms measure “how many times you can divide/multiply” — core to trees, heaps, binary search, and divide-and-conquer.
- **When this appears:** Balanced BST height, heap indexing, binary search iterations, master theorem exponents.

### Identities

```
log(ab)     = log a + log b
log(a/b)    = log a - log b
log(aᵇ)     = b log a
a^(log_b n) = n           (when b^(log_b n) = n)
log_b a     = 1 / log_a b
n = 2^h  ⟺  h = log₂ n    (tree height)
```

### Integer logarithms

- `⌊log₂ n⌋ + 1` = number of bits to represent `n` (for `n > 0`)
- Height of binary heap with `n` nodes: `⌊log₂ n⌋`
- Number of levels in binary search on `[0, n)`: `⌈log₂(n+1)⌉`

### Exponentiation by squaring

Compute `aⁿ` in **O(log n)** multiplications:

```
power(a, n):
  if n == 0: return 1
  if n even: return power(a*a, n/2)
  else:      return a * power(a*a, n/2)
```

With modulo: apply `% MOD` after each multiply (modular exponentiation).

---

## 04 — Recurrence Relations

- **Core idea:** Express total cost `T(n)` in terms of smaller subproblems; solve to get asymptotic bound.
- **When this appears:** Divide-and-conquer, recursion, DP with clear substructure, tree traversals.

### Common recurrences and solutions

| Recurrence | Solution | Example |
|------------|----------|---------|
| `T(n) = T(n-1) + O(1)` | O(n) | Linear recursion |
| `T(n) = T(n-1) + O(n)` | O(n²) | Naive recursive sum |
| `T(n) = T(n/2) + O(1)` | O(log n) | Binary search |
| `T(n) = T(n/2) + O(n)` | O(n) | Partition-style (one half) |
| `T(n) = 2T(n/2) + O(1)` | O(n) | Tree traversal all nodes |
| `T(n) = 2T(n/2) + O(n)` | O(n log n) | Merge sort |
| `T(n) = T(n/2) + O(n log n)` | O(n log n) | |
| `T(n) = T(n-1) + O(log n)` | O(n log n) | |
| `T(n) = T(n-1) + O(2ⁿ)` | O(2ⁿ) | Naive Fibonacci |
| `T(n) = T(n-1) + T(n-2)` | O(φⁿ), φ=(1+√5)/2 | Fibonacci |
| `T(n) = 3T(n/2) + O(n)` | O(n^log₂ 3) ≈ O(n^1.585) | Karatsuba-style |

### Substitution method

1. Guess form (often from recursion tree).
2. Prove by induction with constants.

### Recursion tree method

- Count work per level; sum over `log_b n` levels.
- Example: `2T(n/2) + cn` → levels sum to `cn log₂ n`.

### Master theorem (CLRS)

For `T(n) = aT(n/b) + f(n)`, `a ≥ 1`, `b > 1`, let `n^(log_b a)` be the “leaf cost.”

1. If `f(n) = O(n^(log_b a - ε))` for some `ε > 0` → **T(n) = Θ(n^(log_b a))**
2. If `f(n) = Θ(n^(log_b a) logᵏ n)` → **T(n) = Θ(n^(log_b a) log^(k+1) n)**
3. If `f(n) = Ω(n^(log_b a + ε))` and `af(n/b) ≤ cf(n)` for some `c < 1` → **T(n) = Θ(f(n))**

**Examples:**
- Merge sort: `a=2, b=2, f(n)=Θ(n)` → case 2 with `k=0` → **Θ(n log n)**
- Binary search: `a=1, b=2, f(n)=Θ(1)` → **Θ(log n)**
- Strassen (conceptually): `a=7, b=2` → **Θ(n^log₂ 7)**

### Akra–Bazzi (generalization)

For `T(n) = Σᵢ aᵢ T(n/bᵢ) + g(n)` with regularity conditions:

`T(n) = Θ(nᵖ (1 + ∫₁ⁿ g(u)/u^(p+1) du))` where `p` satisfies `Σᵢ aᵢ/bᵢᵖ = 1`.

Use when subproblem sizes differ (e.g. randomized quicksort).

---

## 05 — Number Theory

- **Core idea:** Integer structure (divisibility, primes, residues) powers hashing, crypto-style problems, GCD/LCM puzzles, and modular DP.
- **When this appears:** “Divisible by k”, coprime pairs, cycle detection in modular arithmetic, combinatorics with overflow.

### Divisibility

- `a | b` means `b = ka` for some integer `k`.
- If `a | b` and `a | c` then `a | (bx + cy)` for integers `x, y`.
- **Euclidean algorithm (GCD):**  
  `gcd(a, b) = gcd(b, a mod b)`, base `gcd(a, 0) = a`.  
  Runs in **O(log min(a,b))** divisions.
- **Extended Euclidean:** finds `x, y` such that `ax + by = gcd(a,b)`.
- **LCM:** `lcm(a, b) = |a·b| / gcd(a, b)`.

### Modular arithmetic

All operations mod `m` (residue ring **Z_m**):

```
(a + b) mod m = ((a mod m) + (b mod m)) mod m
(a - b) mod m = ((a mod m) - (b mod m) + m) mod m
(a · b) mod m = ((a mod m) · (b mod m)) mod m
```

**Do not** distribute modulo over division blindly: use **modular inverse**.

**Modular inverse** of `a` mod `m` exists iff `gcd(a, m) = 1`:

`a⁻¹ ≡ x (mod m)` where `ax ≡ 1 (mod m)` (via extended Euclidean).

**Fermat’s little theorem** (p prime, a not divisible by p):

`a^(p-1) ≡ 1 (mod p)`  ⇒  `a⁻¹ ≡ a^(p-2) (mod p)`

**Euler’s theorem** (generalization): if `gcd(a,n)=1`, then `a^φ(n) ≡ 1 (mod n)`.

### φ(n) — Euler totient

`φ(n)` = count of integers in `[1, n]` coprime to `n`.

- If `p` prime: `φ(p) = p - 1`
- If `p` prime: `φ(pᵏ) = pᵏ - p^(k-1)`
- Multiplicative: if `gcd(a,b)=1` then `φ(ab) = φ(a)φ(b)`

### Primes

- **Fundamental theorem:** every integer `> 1` factors uniquely into primes.
- **Sieve of Eratosthenes:** all primes ≤ `n` in **O(n log log n)** time.
- **Trial division** up to `√n` tests primality in **O(√n)**.
- **Prime number theorem:** π(n) ≈ n / ln n (primes near `n` roughly one per `ln n`).

### Chinese Remainder Theorem (CRT)

If moduli `m₁, …, m_k` are pairwise coprime, the system `x ≡ aᵢ (mod mᵢ)` has a unique solution mod `M = Π mᵢ`.

Useful for big-number modular computations split into coprime factors.

### Useful congruences

- `a ≡ b (mod m)` ⟺ `m | (a - b)`
- Sliding window with modulo: subtract `old * factor` before adding new (rolling hash).

---

## 06 — Combinatorics & Counting

- **Core idea:** Count arrangements, subsets, and paths without enumerating — often via formulas, recurrences, or inclusion–exclusion.
- **When this appears:** Backtracking pruning, DP counting, probability, “how many ways” problems.

### Basic counting

- **Multiplication rule:** independent choices multiply.
- **Addition rule:** disjoint cases add.
- **Pigeonhole principle:** `n+1` items in `n` bins ⇒ some bin has ≥ 2.

### Permutations & combinations

| Object | Formula | Notes |
|--------|---------|-------|
| Permutations of `n` distinct | `n!` | Order matters |
| `k`-permutations of `n` | `P(n,k) = n!/(n-k)!` | No repeat |
| Combinations (binomial) | `C(n,k) = n! / (k!(n-k)!)` | Order irrelevant |
| Multiset permutations | `n! / (n₁! n₂! …)` | Repeated symbols |
| Circular permutations | `(n-1)!` | Fixed rotation equivalence |

**Pascal’s identity:** `C(n,k) = C(n-1,k-1) + C(n-1,k)`

**Binomial theorem:** `(x + y)ⁿ = Σₖ C(n,k) xᵏ y^(n-k)`

### Computing C(n,k) in code

- **O(k)** with multiplicative formula and early exit when `k > n/2` (use symmetry `C(n,k)=C(n,n-k)`).
- **O(n²)** Pascal table for many queries.
- **Modulo prime p:** use Fermat `C(n,k) ≡ n! · (k!)^(-1) · ((n-k)!)^(-1) (mod p)` with precomputed factorials.

### Catalan numbers

`C_n = (1/(n+1)) C(2n, n) = C(2n,n) - C(2n, n-1)`

Counts: valid BSTs of `n` keys, balanced parenthesis strings, non-crossing partitions, many grid paths.

Recurrence: `C₀=1`, `C_{n+1} = Σᵢ₌₀ⁿ Cᵢ C_{n-i}`.

First values: 1, 1, 2, 5, 14, 42, 132, …

### Stirling numbers (awareness)

- **1st kind** (cycles/permutations): permutations of `n` with `k` cycles.
- **2nd kind** (partitions): partition `n` elements into `k` non-empty subsets.

Appear in advanced counting and DP.

### Inclusion–exclusion

`|A₁ ∪ … ∪ Aₙ| = Σ|Aᵢ| - Σ|Aᵢ∩Aⱼ| + … + (-1)^(n+1)|A₁∩…∩Aₙ|`

Use when counting “at least one” / forbidden patterns.

### Stars and bars

Non-negative integer solutions to `x₁ + … + x_k = n`: **C(n+k-1, k-1)**.

Positive integer solutions: **C(n-1, k-1)**.

---

## 07 — Probability & Expected Value

- **Core idea:** Analyze randomized algorithms and average-case behavior; linearity of expectation simplifies many “count expected X” problems.
- **When this appears:** Quickselect, randomized quicksort, reservoir sampling, hashing collision analysis, Monte Carlo.

### Axioms & rules

- `0 ≤ P(A) ≤ 1`, `P(Ω) = 1`
- `P(A ∪ B) = P(A) + P(B) - P(A ∩ B)`
- **Conditional:** `P(A|B) = P(A∩B)/P(B)`
- **Independence:** `P(A∩B) = P(A)P(B)`

### Expectation

**Linearity (always, even if dependent):**

`E[X + Y] = E[X] + E[Y]`  
`E[cX] = c E[X]`

**Indicator variables:** `X = Σ Iᵢ` where `Iᵢ = 1` if event `i` occurs.

`E[X] = Σ P(event i)` — powerful for counting expected occurrences (e.g. fixed points, inversions in random perm).

**Markov’s inequality:** `P(X ≥ t) ≤ E[X]/t` (for non-negative X).

**Chebyshev:** `P(|X - μ| ≥ kσ) ≤ 1/k²`.

### Classic examples

| Problem | Result |
|---------|--------|
| Random quicksort (average) | **O(n log n)** comparisons |
| Quickselect average | **O(n)** |
| Birthday paradox | 23 people ⇒ >50% chance shared birthday |
| Coupon collector | **E = n H(n) ≈ n ln n + γn** trials to collect all `n` coupons |
| Hash table load α | Expected probes (chaining) **O(1 + α)** |

---

## 08 — Graph Theory (Structural Math)

- **Core idea:** Graphs have invariant counts and bounds that explain algorithm limits and correctness.
- **When this appears:** Trees, connectivity, planarity, coloring bounds, network flow caps.

### Basic definitions

- `G = (V, E)`, `n = |V|`, `m = |E|`
- **Handshaking lemma:** `Σᵥ deg(v) = 2m`
- **Maximum edges** (simple undirected): `m ≤ n(n-1)/2`
- **Maximum edges** (simple directed): `m ≤ n(n-1)`

### Trees

- Connected acyclic graph.
- **`|E| = |V| - 1`** (and connected ⇒ acyclic iff this holds).
- Unique simple path between any two vertices.
- Adding an edge creates exactly one cycle; removing any edge disconnects.

### Spanning tree

- `n - 1` edges connect all `n` vertices.
- Minimum spanning tree algorithms (Kruskal, Prim) rely on cut/greedy properties.

### Euler trails & circuits

- Euler circuit exists iff graph connected and **all vertices even degree**.
- Euler trail (not circuit) iff connected and **exactly 0 or 2 odd-degree vertices**.

### Planar graphs (Euler formula)

For connected planar embedding: **`v - e + f = 2`** (faces `f` includes outer face).

⇒ simple planar graph: **`e ≤ 3v - 6`** (for `v ≥ 3`).

### Bipartite graphs

- No odd cycles.
- 2-colorable iff bipartite.
- **Kőnig:** in bipartite graph, size of max matching = size of min vertex cover.

### Distance metrics

- **Diameter:** max shortest-path distance between any pair.
- **Radius:** min over vertices of max distance to others.
- In unweighted graph, BFS from each vertex finds eccentricities in **O(V(V+E))**; can optimize with APSP if needed.

---

## 09 — Trees & Heaps (Structural)

- **Core idea:** Hierarchical structure implies logarithmic depth when balanced.
- **When this appears:** BSTs, heaps, segment trees, Fenwick trees, trie depth.

### Binary tree counts

- **Catalan `C_n`** = number of structurally distinct binary trees with `n` nodes.
- Full binary tree with `L` leaves: **internal nodes = L - 1**, total nodes = `2L - 1`.
- Complete binary tree with `n` nodes: height **⌊log₂ n⌋**.

### Heap array indexing (0-based)

| Relation | Formula |
|----------|---------|
| Parent of `i` | `(i - 1) / 2` |
| Left child | `2i + 1` |
| Right child | `2i + 2` |
| Level of node `i` | `⌊log₂(i+1)⌋` |

- Build heap from array: **O(n)** (bottom-up sift), not O(n log n).
- Extract-min `n` times: **O(n log n)**.

### BST average vs worst

- Random insert order: expected height **O(log n)**.
- Sorted insert: height **O(n)** (degenerate).

### Trie

- String set of total length `N` over alphabet `σ`: storage **O(N)**, lookup **O(L)** for string length `L`.

---

## 10 — Dynamic Programming — Mathematical Structure

- **Core idea:** Optimal substructure + overlapping subproblems; often a recurrence with clear state.
- **When this appears:** Optimization, counting paths, knapsack, LIS, edit distance.

### Fibonacci

Recurrence: `F(n) = F(n-1) + F(n-2)`, `F(0)=0`, `F(1)=1`.

- Naive recursion: **O(φⁿ)**
- DP / iterative: **O(n)** time, **O(1)** space
- Matrix form: `[F(n+1), F(n)]ᵀ = [[1,1],[1,0]]ⁿ [1,0]ᵀ` → **O(log n)** via fast matrix pow

**Binet (closed form):** `F(n) = (φⁿ - ψⁿ)/√5`, φ=(1+√5)/2, ψ=(1-√5)/2.

### Knapsack variants

| Variant | Typical complexity |
|---------|-------------------|
| 0/1 knapsack | O(nW) pseudo-polynomial (W = capacity) |
| Unbounded knapsack | O(nW) |
| Fractional knapsack | O(n log n) greedy |

**Note:** NP-hard in `W` encoded in bits; pseudo-polynomial means polynomial in numeric value, not input size.

### Longest increasing subsequence (LIS)

- DP: **O(n²)**
- Patience sorting + binary search: **O(n log n)**

### Edit distance (Levenshtein)

Recurrence on prefixes; **O(mn)** time and space (optimizable to O(min(m,n)) space).

### DAG shortest/longest path

**O(V + E)** with topological order — no negative cycles in DAG.

---

## 11 — Greedy — Exchange Arguments & Matroids (Awareness)

- **Core idea:** Prove greedy choice is safe via exchange argument or matroid structure.
- **When this appears:** Interval scheduling, Huffman coding, MST, fractional knapsack.

### Interval scheduling

Sort by finish time; greedy picks maximum non-overlapping set — optimal by exchange argument.

### Huffman coding

Build min binary tree by merging two least frequencies — minimizes weighted code length.

### Matroid greedy (CLRS)

If feasible sets form a matroid, greedy by weight is optimal (e.g. MST graphic matroid).

---

## 12 — Bit Manipulation & Binary Representations

- **Core idea:** Integers are bit vectors; bitwise ops encode set membership, parity, and powers of two in O(1).
- **When this appears:** Subset enumeration, XOR tricks, flags, low-level optimizations.

### Operators

| Op | Meaning |
|----|---------|
| `a & b` | Bitwise AND |
| `a \| b` | Bitwise OR |
| `a ^ b` | XOR (xor cancel: `a^a=0`) |
| `~a` | NOT |
| `a << k` | Left shift (= multiply by `2ᵏ`) |
| `a >> k` | Right shift (signed: implementation-defined for negative) |

### Classic identities

```
n & (n - 1)     clears lowest set bit
n & (-n)        isolates lowest set bit (two's complement)
2^k             == 1 << k
n power of 2    ⟺ n > 0 && (n & (n-1)) == 0
popcount(n)     # of 1-bits (builtin: __builtin_popcount)
```

### XOR properties

- Commutative, associative: order doesn’t matter.
- `a ^ 0 = a`, `a ^ a = 0`
- Find single non-duplicate in pairs: XOR all elements.

### Subset iteration

For mask `0..(1<<n)-1`, each mask is a subset of `n` items — **O(2ⁿ)**.

Iterate submasks: `for (sub = mask; sub; sub = (sub-1) & mask)` — **O(3ⁿ)** total over all masks (use carefully).

---

## 13 — Geometry (Competitive / Interview)

- **Core idea:** Vectors, cross products, and distances answer orientation, intersection, and containment.
- **When this appears:** Convex hull, line intersection, closest pair, grid geometry.

### 2D vectors

- Point: `(x, y)`. Vector: `B - A`.
- **Distance:** `√((x₂-x₁)² + (y₂-y₁)²)` — compare squared distance to avoid sqrt.
- **Dot product:** `a·b = |a||b|cos θ = a_x b_x + a_y b_y`
  - `a·b = 0` ⟹ perpendicular.
- **Cross product (2D scalar):** `a × b = a_x b_y - a_y b_x`
  - Sign gives orientation (CCW > 0, CW < 0, collinear = 0).
  - `|a × b| = |a||b|sin θ` = area of parallelogram.

### Orientation test

For three points A, B, C: `orient(A,B,C) = (B-A) × (C-A)`.

### Line & segment

- Parametric line: `P = P₀ + t(P₁ - P₀)`.
- Segment intersection uses orientation tests on both segments.

### Polygon

- **Area (shoelace):** `A = ½|Σ(xᵢ yᵢ₊₁ - xᵢ₊₁ yᵢ)|`
- Point-in-polygon: ray casting or winding number.

### Convex hull

- Graham scan / Andrew monotone chain: **O(n log n)**.
- Gift wrapping (Jarvis march): **O(nh)** (h = hull size).

### Grid geometry

- Manhattan distance: `|x₁-x₂| + |y₁-y₂|`
- Cells at Manhattan distance `d`: diamond shape, count grows as **O(d²)**

---

## 14 — String Algorithms — Math Behind the Patterns

- **Core idea:** Prefix functions and rolling hashes use modular arithmetic and recurrence on prefixes.
- **When this appears:** Pattern matching, duplicate detection, palindrome centers, substring search.

### Rolling hash

`H(s[0..i]) = (H(s[0..i-1]) · B + s[i]) mod M`

Update sliding window in **O(1)** with add/remove and precomputed `B^(len) mod M`.

Choose large prime `M`, base `B`; watch collisions (double hash if needed).

### KMP prefix function π

`π[i]` = length of longest proper prefix of `s[0..i]` that is also suffix.

Computed in **O(n)**; search in **O(n + m)**.

### Z-algorithm

`Z[i]` = LCP of `s` and `s[i..]`. Also **O(n)**.

### Suffix structures (awareness)

- Suffix array + LCP: powerful substring queries.
- Suffix tree/automaton: linear construction (Ukkonen).

---

## 15 — Numerical Stability & Integer Overflow

- **Core idea:** Language ints wrap; use wider types or modular arithmetic before intermediate products overflow.
- **When this appears:** Combinatorics, graph weights, matrix problems, products of coordinates.

### Safe practices

- Check before `a * b`: if `a > MAX / b` then overflow.
- Use **64-bit** (`long long`) for sums/products in many contest bounds.
- Modular multiplication: `(a * b) % m` with `a,b < m` may need `((a%m)*(b%m))%m` in 128-bit or `mulmod`.
- **Prefix sums** on large arrays: watch cumulative overflow.

### Modular subtraction

Always add `m` before final `% m` when subtracting: `(a - b + m) % m`.

---

## 16 — Matrices & Fast Exponentiation (Advanced)

- **Core idea:** Linear recurrences and graph reachability reduce to matrix multiplication.
- **When this appears:** Fibonacci in O(log n), Markov chains, counting paths of length `k`, DP on small state spaces.

### Matrix multiplication

`(AB)_{ij} = Σ_k A_{ik} B_{kj}`

Naive: **O(n³)** for `n×n` matrices; Strassen **O(n^2.807)** (theory).

### Path counting

Adjacency matrix `A`: `(A^k)_{ij}` = number of paths of length `k` from `i` to `j`.

### Linear recurrence

If `F(n)` depends on previous `k` values linearly, use `k×k` transition matrix raised to `(n-k)`.

---

## 17 — Information Theory & Entropy (Awareness)

- **Core idea:** Shannon entropy lower-bounds average code length; explains why comparison sorts need Ω(n log n) bits of information to distinguish `n!` permutations.
- **When this appears:** Compression (Huffman), lower bounds, randomized bit complexity.

**Entropy:** `H(X) = -Σ p(x) log₂ p(x)` bits.

For uniform distribution over `n` outcomes: `H = log₂ n`.

---

## 18 — Order Statistics

- **Core idea:** Find the k-th smallest without full sort.
- **When this appears:** Median, quickselect, k-th largest in stream (with heap).

### Quickselect

Average **O(n)**, worst **O(n²)** (mitigated by random pivot).

### Median of medians

Deterministic **O(n)** worst-case selection (theory; rarely implemented in interviews).

### Heap method for k-th largest

Min-heap of size `k`: **O(n log k)** time, **O(k)** space.

---

## Fast Math Picker

Use this when stuck on the numeric side of a problem:

| When you see… | Math tool… |
|---------------|------------|
| Nested loops / recursion depth | Summations, Master theorem, recursion tree |
| “How many ways?” | Combinations, Catalan, inclusion–exclusion, DP count |
| “Divisible by k” / huge integers | Modular arithmetic, GCD, inverse, CRT |
| Sorted search / halving | Logarithms, O(log n) recurrence |
| Tree / heap depth | `log₂ n`, heap index formulas |
| Randomized algorithm average case | Linearity of expectation, indicator variables |
| XOR / subset / power of 2 | Bit tricks, masks `1<<k` |
| Geometry turn / intersection | Cross product orientation |
| String pattern match | Rolling hash, KMP π, Z-array |
| Linear recurrence for huge `n` | Matrix exponentiation |
| Overflow in products | 64-bit, modular mul, check before multiply |
| Sorting lower bound | `n!` leaves ⇒ Ω(n log n) comparisons |

---

## Formula Sheet (one-page recap)

```
Arithmetic sum:     Σ i   = n(n+1)/2
Geometric sum:      Σ rⁱ  = (r^(n+1)-1)/(r-1)
Harmonic:           H(n)  ≈ ln n + γ

GCD:                gcd(a,b) = gcd(b, a mod b)
LCM:                lcm(a,b) = a·b / gcd(a,b)

Binomial:           C(n,k) = n! / (k!(n-k)!)
Catalan:            C_n = C(2n,n)/(n+1)

Master:             T(n)=aT(n/b)+f(n) → compare f(n) to n^(log_b a)

Tree edges:         |E| = |V| - 1
Handshaking:        Σ deg = 2|E|

Heap parent:        (i-1)/2;  children: 2i+1, 2i+2
Binary search:      ≤ ⌈log₂(n+1)⌉ steps

Mod add/mul:        (a+b)%m, (a*b)%m;  inv when gcd(a,m)=1
Fermat inverse:     a^(p-2) mod p  (p prime)

2D cross:           (b-a) × (c-a)  → orientation sign
Shoelace area:      ½|Σ(x_i y_{i+1} - x_{i+1} y_i)|
```

---

## How to Study This File

1. Pick one section tied to this week’s topic (e.g. graphs → §08, DP → §10).
2. Write the key formulas on paper from memory.
3. Solve one “count / bound / modulo” sub-step of a medium problem using only this sheet.
4. Cross-link with `DRILL_CONCEPTS.md` for pattern triggers and `study_code/.../05_complexity_glance` for complexity drills.
5. Log one formula you forgot and where it would have saved time.

**Related docs:** `DRILL_CONCEPTS.md` (patterns), `STUDY_PLAN.md` (schedule), `READING_PATTERNS.md` (complexity pass in code reading).
