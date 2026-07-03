/**
 * REFLEX DRILL 04 — Binary Search
 *
 * RUN: node study_play/drills/04_binary_search_reflex.js
 */

// TODO: REFLEX — return index of target or -1
function binarySearch(nums, target) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — insert position (first index where nums[i] >= target)
function searchInsert(nums, target) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — find minimum in rotated sorted array (no duplicates)
function findMinRotated(nums) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — is target in nums (classic template with lo <= hi)
function isTargetPresent(nums, target) {
  throw new Error("Implement from memory");
}

function assert(name, cond) {
  if (!cond) throw new Error(`FAIL: ${name}`);
  console.log(`PASS: ${name}`);
}

assert("binarySearch found", binarySearch([-1, 0, 3, 5, 9, 12], 9) === 4);
assert("binarySearch missing", binarySearch([-1, 0, 3, 5, 9, 12], 2) === -1);
assert("searchInsert exist", searchInsert([1, 3, 5, 6], 5) === 2);
assert("searchInsert new", searchInsert([1, 3, 5, 6], 2) === 1);
assert("findMinRotated", findMinRotated([4, 5, 6, 7, 0, 1, 2]) === 0);
assert("isTargetPresent", isTargetPresent([1, 2, 3, 4, 5], 3) === true);

console.log("\nAll binary search reflex drills passed.");

module.exports = { binarySearch, searchInsert, findMinRotated, isTargetPresent };
