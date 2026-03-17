// given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.

// your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

// example 1:

// input: nums = [1,2,2,3,1]
// output: 2
// explanation:
// the input array has a degree of 2 because both elements 1 and 2 appear twice.
// of the subarrays that have the same degree:
// [1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
// the shortest length is 2. so return 2.

// example 2:

// input: nums = [1,2,2,3,1,4,2]
// output: 6
// explanation:
// the degree is 3 because the element 2 is repeated 3 times.
// so [2,2,3,1,4,2] is the shortest subarray, therefore returning 6.

// constraints:

//     nums.length will be between 1 and 50,000.
//     nums[i] will be an integer between 0 and 49,999.

/**
 * @param {number[]} nums
 * @return {number}
 */
var findshortestsubarray = function (nums) {
  const counts = new Map();
  const first = new Map();
  const last = new Map();
  let maxDegree = 0;

  for (let i = 0; i < nums.length; i++) {
    const num = nums[i];

    if (!first.has(num)) {
      first.set(num, i);
    }

    last.set(num, i);

    const newCount = (counts.get(num) | 0) + 1;
    counts.set(num, newCount);
    maxDegree = Math.max(maxDegree, newCount);
  }

  let minLength = nums.length;

  for (let [num, count] of counts) {
    if (count === maxDegree) {
      const span = last.get(num) - first.get(num) + 1;
      minLength = Math.min(minLength, span);
    }
  }
  return minLength;
};

console.log(findshortestsubarray([1, 2, 2, 3, 1]));
console.log(findshortestsubarray([1, 2, 2, 3, 1, 4, 2]));
