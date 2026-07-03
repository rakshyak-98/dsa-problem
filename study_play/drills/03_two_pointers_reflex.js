/**
 * REFLEX DRILL 03 — Two Pointers & Sliding Window
 *
 * RUN: node study_play/drills/03_two_pointers_reflex.js
 */

// TODO: REFLEX — remove duplicates from sorted array, return new length
function removeDuplicates(nums) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — move all zeroes to end in-place
function moveZeroes(nums) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — max area container (heights array)
function maxArea(heights) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — is palindrome (alphanumeric only, ignore case)
function isPalindrome(s) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — max sum of subarray of size k (sliding window fixed)
function maxSumSubarrayK(nums, k) {
  throw new Error("Implement from memory");
}

function assert(name, cond) {
  if (!cond) throw new Error(`FAIL: ${name}`);
  console.log(`PASS: ${name}`);
}

const dup = [1, 1, 2, 2, 3];
assert("removeDuplicates len", removeDuplicates(dup) === 3);
assert("removeDuplicates vals", JSON.stringify(dup.slice(0, 3)) === "[1,2,3]");

const zeros = [0, 1, 0, 3, 12];
moveZeroes(zeros);
assert("moveZeroes", JSON.stringify(zeros) === "[1,3,12,0,0]");

assert("maxArea", maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]) === 49);
assert("isPalindrome", isPalindrome("A man, a plan, a canal: Panama") === true);
assert("maxSumSubarrayK", maxSumSubarrayK([2, 1, 5, 1, 3, 2], 3) === 9);

console.log("\nAll two-pointer reflex drills passed.");

module.exports = { removeDuplicates, moveZeroes, maxArea, isPalindrome, maxSumSubarrayK };
