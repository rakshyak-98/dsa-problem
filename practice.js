/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
  let lookup = new Map();
  for (let i = 0; i < nums.length; i++) {
    const diff = target - nums[i];
    if (lookup.has(diff)) return [lookup.get(diff), i];
    lookup.set(nums[i], i);
  }
  return [];
};

console.log(twoSum([2, 7, 11, 15], 9));
console.log(twoSum([3, 2, 4], 6));
console.log(twoSum([3, 3], 6));
