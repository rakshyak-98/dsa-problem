// You are given a positive integer n.

// Let s be the binary representation of n without leading zeros.

// The reverse of a binary string s is obtained by writing the characters of s in the opposite order.

// You may flip any bit in s (change 0 → 1 or 1 → 0). Each flip affects exactly one bit.

// Return the minimum number of flips required to make s equal to the reverse of its original form.

// Example 1:

// Input: n = 7

// Output: 0

// Explanation:

// The binary representation of 7 is "111". Its reverse is also "111", which is the same. Hence, no flips are needed.

// Example 2:

// Input: n = 10

// Output: 4

// Explanation:

// The binary representation of 10 is "1010". Its reverse is "0101". All four bits must be flipped to make them equal. Thus, the minimum number of flips required is 4.

// Constraints:

//     1 <= n <= 109

/**
 * @param {number} n
 * @return {number}
 */
var minimumFlips = function (n) {
  const s = n.toString(2); 
  const reversed = s.split('').reverse().join('');

  let flips = 0;
  for(let i = 0; i < s.length; i++){
    if(s[i] !== reversed[i]){
     flips++; 
    }
  }
  return flips;
};

console.log(minimumFlips(7));
console.log(minimumFlips(10));
