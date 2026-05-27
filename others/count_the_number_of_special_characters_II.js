// You are given a string word. A letter c is called special if it appears both in lowercase and uppercase in word, and every lowercase occurrence of c appears before the first uppercase occurrence of c.

// Return the number of special letters in word.

// Example 1:

// Input: word = "aaAbcBC"

// Output: 3

// Explanation:

// The special characters are 'a', 'b', and 'c'.

// Example 2:

// Input: word = "abc"

// Output: 0

// Explanation:

// There are no special characters in word.

// Example 3:

// Input: word = "AbBCab"

// Output: 0

// Explanation:

// There are no special characters in word.

// Constraints:

//     1 <= word.length <= 2 * 105
//     word consists of only lowercase and uppercase English letters.

/**
 * @param {string} word
 * @return {number}
 */
var numberOfSpecialChars = function (word) {
  const lower = new Array(26).fill(-1);
  const upper = new Array(26).fill(-1);

  for (let i = 0; i < word.length; i++) {
    const charCode = word.charCodeAt(i);

    if (charCode >= 97 && charCode <= 122) {
      lower[charCode - 97] = i;
    } else if (charCode >= 65 && charCode <= 90) {
      const index = charCode - 65;

      if (upper[index] == -1) {
        upper[index] = i;
      }
    }
  }

  let specialCount = 0;

  for (let i = 0; i < 26; i++) {
    if (lower[i] !== -1 && upper[i] !== -1 && lower[i] < upper[i]) {
      specialCount++;
    }
  }
  return specialCount;
};

console.log(numberOfSpecialChars("aaAbcBC"));
console.log(numberOfSpecialChars("abc"));
