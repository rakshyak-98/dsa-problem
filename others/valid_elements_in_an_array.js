// You are given an integer array nums.

// An element nums[i] is considered valid if it satisfies at least one of the following conditions:

//     It is strictly greater than every element to its left.
//     It is strictly greater than every element to its right.

// The first and last elements are always valid.

// Return an array of all valid elements in the same order as they appear in nums.

// Example 1:

// Input: nums = [1,2,4,2,3,2]

// Output: [1,2,4,3,2]

// Explanation:

//     nums[0] and nums[5] are always valid.
//     nums[1] and nums[2] are strictly greater than every element to their left.
//     nums[4] is strictly greater than every element to its right.
//     Thus, the answer is [1, 2, 4, 3, 2].

// Example 2:

// Input: nums = [5,5,5,5]

// Output: [5,5]

// Explanation:

//     The first and last elements are always valid.
//     No other elements are strictly greater than all elements to their left or to their right.
//     Thus, the answer is [5, 5].

// Example 3:

// Input: nums = [1]

// Output: [1]

// Explanation:

// Since there is only one element, it is always valid. Thus, the answer is [1].

// Constraints:

//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100

/**
 * @param {number[]} nums
 * @return {number[]}
 */
var findValidElements = function (nums) {
  let n = nums.length;

  if (n == 0) return false;
  if (n == 1) return [nums[0]];

  const isValid = new Array(n).fill(false);
  isValid[0] = true;
  isValid[n - 1] = true;

  let leftMax = nums[0];
  for (let i = 1; i < n - 1; i++) {
    if (nums[i] > leftMax) {
      isValid[i] = true;
      leftMax = nums[i];
    }
  }

  let rightMax = nums[n - 1];
  for (let i = n - 2; i > 0; i--) {
    if (nums[i] > rightMax) {
      isValid[i] = true;
      rightMax = nums[i];
    }
  }

  const result = [];
  for (let i = 0; i < n; i++) {
    if (isValid[i]) {
      result.push(nums[i]);
    }
  }

  return result;
};

console.log(findValidElements([1, 2, 4, 2, 3, 2]));
console.log(findValidElements([5, 5, 5, 5]));
console.log(findValidElements([1]));
