// Given an integer array nums, find three numbers whose product is maximum and return the maximum product.

// Example 1:

// Input: nums = [1,2,3]
// Output: 6

// Example 2:

// Input: nums = [1,2,3,4]
// Output: 24

// Example 3:

// Input: nums = [-1,-2,-3]
// Output: -6

// Constraints:

//     3 <= nums.length <= 104
//     -1000 <= nums[i] <= 1000

/**
 * @param {number[]} nums
 * @return {number}
 */
var maximumProduct = function (nums) {
  let max1 = -Infinity,
    max2 = -Infinity,
    max3 = -Infinity,
    min1 = Infinity,
    min2 = Infinity;
  for (let n of nums) {
    if (n > max1) {
      max3 = max2;
      max2 = max1;
      max1 = n;
    } else if (n > max2) {
      max3 = max2;
      max2 = n;
    } else if (n > max3) {
      max3 = n;
    }

    if (n < min1) {
      min2 = min1;
      min1 = n;
    } else if (n < min2) {
      min2 = n;
    }
  }
  return Math.max(max1 * max2 * max3, min1 * min2 * max1);
};

console.log(maximumProduct([1, 2, 3]));
console.log(maximumProduct([-1, -2, -3]));
