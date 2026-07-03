/**
 * REFLEX DRILL 02 — Hashing
 *
 * GOAL: Hash map patterns = automatic.
 * RUN: node study_play/drills/02_hashing_reflex.js
 */

// TODO: REFLEX — classic two sum (return indices)
function twoSum(nums, target) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — return true if any value appears twice
function containsDuplicate(nums) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — return Map of value -> frequency
function frequencyMap(arr) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — first character that appears only once; return '' if none
function firstUniqueChar(s) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — group anagrams (return array of groups)
function groupAnagrams(strs) {
  throw new Error("Implement from memory");
}

function assert(name, cond) {
  if (!cond) throw new Error(`FAIL: ${name}`);
  console.log(`PASS: ${name}`);
}

assert("twoSum", JSON.stringify(twoSum([2, 7, 11, 15], 9)) === "[0,1]");
assert("containsDuplicate true", containsDuplicate([1, 2, 3, 1]) === true);
assert("containsDuplicate false", containsDuplicate([1, 2, 3, 4]) === false);

const freq = frequencyMap(["a", "b", "a", "c"]);
assert("frequencyMap", freq.get("a") === 2 && freq.get("b") === 1);

assert("firstUniqueChar", firstUniqueChar("leetcode") === "l");
assert("firstUniqueChar none", firstUniqueChar("aabb") === "");

const groups = groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]);
assert(
  "groupAnagrams count",
  groups.length === 3 && groups.some((g) => g.length === 3 && g.includes("eat"))
);

console.log("\nAll hashing reflex drills passed.");

module.exports = { twoSum, containsDuplicate, frequencyMap, firstUniqueChar, groupAnagrams };
