/**
 * REFLEX DRILL 01 — Arrays
 *
 * GOAL: Write every function from memory in < 3 minutes each.
 * RUN: node study_play/drills/01_arrays_reflex.js
 *
 * After filling TODOs, all tests should pass.
 */

// TODO: REFLEX — reverse array in-place, return the same array
function reverseInPlace(arr) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — return index of max element (first max if ties)
function indexOfMax(arr) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — return sum of all elements
function arraySum(arr) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — rotate right by k steps (use modulo on k)
function rotateRight(arr, k) {
  throw new Error("Implement from memory");
}

// TODO: REFLEX — return new array: running sum (prefix as output)
function runningSum(arr) {
  throw new Error("Implement from memory");
}

// --- Self tests (do not edit until your functions pass) ---
function assert(name, cond) {
  if (!cond) throw new Error(`FAIL: ${name}`);
  console.log(`PASS: ${name}`);
}

assert("reverseInPlace", JSON.stringify(reverseInPlace([1, 2, 3])) === "[3,2,1]");
assert("indexOfMax", indexOfMax([3, 1, 4, 4]) === 2);
assert("arraySum", arraySum([1, 2, 3, 4]) === 10);
assert("rotateRight", JSON.stringify(rotateRight([1, 2, 3, 4, 5], 2)) === "[4,5,1,2,3]");
assert("rotateRight k>len", JSON.stringify(rotateRight([1, 2], 5)) === "[2,1]");
assert("runningSum", JSON.stringify(runningSum([1, 2, 3, 4])) === "[1,3,6,10]");

console.log("\nAll array reflex drills passed.");

module.exports = { reverseInPlace, indexOfMax, arraySum, rotateRight, runningSum };
