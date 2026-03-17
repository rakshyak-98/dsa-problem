// Given a binary array nums, return the maximum number of consecutive 1's in the array.

// Example 1:

// Input: nums = [1,1,0,1,1,1]
// Output: 3
// Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.

// Example 2:

// Input: nums = [1,0,1,1,0,1]
// Output: 2

// Constraints:

//     1 <= nums.length <= 105
//     nums[i] is either 0 or 1.

/**
 * @param {Array} nums
 * @returns Array
 */
var findMaxConsecutiveOnes = function (nums) {
  let max = 0;
  let count = 0;
  let index = [];
  let startIndex = 0;
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] === 1) {
      count++;
    } else {
      count = 0;
      startIndex = i + 1;
    }
    if (count > max) {
      max = count;
      index[0] = startIndex;
      index[1] = i;
    }
  }
  return index;
};

console.log(findMaxConsecutiveOnes([1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1]));
