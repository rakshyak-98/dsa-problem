/**
 * @param {Array} nums
 */

function hasDuplicates(nums) {
  let counter = 0;
  nums.sort((a, b) => a - b);
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] === nums[i - 1]) {
      counter++;
    }
    if (counter > 0) return true;
  }
  return false;
}

console.log(hasDuplicates([1, 2, 3, 4, 3]));
