// You are given a positive integer n and an integer target.

// Return the array of integers of size n such that:

//     The sum of its elements equals target.
//     The absolute values of its elements form a permutation of size n.

// If no such array exists, return an empty array.

// A permutation of size n is a rearrangement of integers 1, 2, ..., n.

// Example 1:

// Input: n = 3, target = 0

// Output: [-3,1,2]

// Explanation:

// The arrays that sum to 0 and whose absolute values form a permutation of size 3 are:

//     [-3, 1, 2]
//     [-3, 2, 1]
//     [-2, -1, 3]
//     [-2, 3, -1]
//     [-1, -2, 3]
//     [-1, 3, -2]
//     [1, -3, 2]
//     [1, 2, -3]
//     [2, -3, 1]
//     [2, 1, -3]
//     [3, -2, -1]
//     [3, -1, -2]

// The lexicographically smallest one is [-3, 1, 2].

// Example 2:

// Input: n = 1, target = 10000000000

// Output: []

// Explanation:

// There are no arrays that sum to 10000000000 and whose absolute values form a permutation of size 1. Therefore, the answer is [].

// Constraints:

//     1 <= n <= 105
//     -1010 <= target <= 1010

/**
 * @param {number} n
 * @param {number} target
 * @return {number[]}
 */
var lexSmallestNegatedPerm = function (n, target) {
  const sumAll = (n * (n + 1)) / 2;
  const diff = sumAll - target;

  if (diff < 0 || diff % 2 !== 0) return [];

  const targetToReduce = diff / 2;
  const flipped = new Set();
  let currentReduce = targetToReduce;

  for (let i = n; i >= 1; i--) {
    if (currentReduce >= i) {
      flipped.add(i);
      currentReduce -= i;
    }
  }

  if (currentReduce !== 0) return [];

  const result = [];
  for (let i = 1; i <= n; i++) {
    if (flipped.has(i)) {
      result.push(-i);
    } else {
      result.push(i);
    }
  }
  return result;
};

console.log(lexSmallestNegatedPerm(1, 0));
console.log(lexSmallestNegatedPerm(1, 10000000000));
