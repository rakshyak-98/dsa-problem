// You are given a 0-indexed binary string s and two integers minJump and maxJump. In the beginning, you are standing at index 0, which is equal to '0'. You can move from index i to index j if the following conditions are fulfilled:

//     i + minJump <= j <= min(i + maxJump, s.length - 1), and
//     s[j] == '0'.

// Return true if you can reach index s.length - 1 in s, or false otherwise.

// Example 1:

// Input: s = "011010", minJump = 2, maxJump = 3
// Output: true
// Explanation:
// In the first step, move from index 0 to index 3.
// In the second step, move from index 3 to index 5.

// Example 2:

// Input: s = "01101110", minJump = 2, maxJump = 3
// Output: false

// Constraints:

//     2 <= s.length <= 105
//     s[i] is either '0' or '1'.
//     s[0] == '0'
//     1 <= minJump <= maxJump < s.length

/**
 * @param {string} s
 * @param {number} minJump
 * @param {number} maxJump
 * @return {boolean}
 */
var canReach = function (s, minJump, maxJump) {
  let n = s.length;
  let arr = new Array(n).fill(0);
  let count = 0;

  if (s[n - 1] == "1") return false;

  for (let i = 0; i < s.length; i++) {
    count += arr[i];

    if (i == 0 || (s[i] == "0" && count > 0)) {
      if (i + minJump < n) arr[i + minJump] += 1;
      if (i + maxJump + 1 < n) arr[i + maxJump + 1] -= 1;
    }
  }
  return count > 0;
};

console.log(canReach("011010", 2, 3));
console.log(canReach("01101110", 2, 3));
