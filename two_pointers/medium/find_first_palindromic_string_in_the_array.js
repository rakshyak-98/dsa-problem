// Given an array of strings words, return the first palindromic string in the array. If there is no such string, return an empty string "".

// A string is palindromic if it reads the same forward and backward.

// Example 1:

// Input: words = ["abc","car","ada","racecar","cool"]
// Output: "ada"
// Explanation: The first string that is palindromic is "ada".
// Note that "racecar" is also palindromic, but it is not the first.

// Example 2:

// Input: words = ["notapalindrome","racecar"]
// Output: "racecar"
// Explanation: The first and only string that is palindromic is "racecar".

// Example 3:

// Input: words = ["def","ghi"]
// Output: ""
// Explanation: There are no palindromic strings, so the empty string is returned.

// Constraints:

//     1 <= words.length <= 100
//     1 <= words[i].length <= 100
//     words[i] consists only of lowercase English letters.

/**
 * @param {string[]} words
 * @return {string}
 */
var firstPalindrome = function (words) {
  for (const word of words) {
    if (isPalindrome(word)) {
      return word;
    }
  }
  return "";
};

function isPalindrome(s) {
  let left = 0;
  let right = s.length - 1;

  while (left < right) {
    if (s[left] !== s[right]) {
      return false;
    }
    left++;
    right--;
  }
  return true;
}

console.log(firstPalindrome(["abc", "car", "ada", "racecar", "cool"]));
console.log(firstPalindrome(["notapalindrome", "racecar"]));
console.log(firstPalindrome(["def", "ghi"]));
