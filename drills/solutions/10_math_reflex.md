# Drill 10 — Math

## gcd
- **Trigger:** common divisor, simplify fractions, coprime checks
- **Pattern:** `while b != 0: a, b = b, a % b`
- **Bug:** forgetting `abs` on negative inputs

## lcm
- **Trigger:** align cycles, merge repeating events
- **Pattern:** `|a*b| / gcd(a,b)` — divide before multiply when possible
- **Bug:** multiplying before dividing → overflow

## modPow
- **Trigger:** huge exponent mod prime (inverse, counting mod m)
- **Pattern:** square-and-multiply; `% mod` every step
- **Bug:** `mod == 1` → answer is always 0

## nCr
- **Trigger:** count combinations, Pascal paths
- **Pattern:** `k = min(k, n-k)` then multiplicative formula
- **Bug:** factorial overflow — use iterative multiply/divide

## isPrime
- **Trigger:** prime sieve building block, factorization stop at √n
- **Pattern:** test 2, then odd divisors to `i*i <= n`
- **Bug:** treating 1 as prime

## powOfTwo
- **Trigger:** bitmask sizing, heap levels, binary tree nodes
- **Pattern:** `n > 0 && (n & (n-1)) == 0`
- **Bug:** forgetting `n > 0` (0 is not a power of 2)
