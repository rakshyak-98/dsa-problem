// Given an array of strings strs, group the together. You can return the answer in any order.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]

// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// Explanation:

//     There is no string in strs that can be rearranged to form "bat".
//     The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
//     The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.

// Example 2:

// Input: strs = [""]

// Output: [[""]]

// Example 3:

// Input: strs = ["a"]

// Output: [["a"]]

// Constraints:

//     1 <= strs.length <= 104
//     0 <= strs[i].length <= 100
//     strs[i] consists of lowercase English letters.

/**
 * @param {string[]} strs
 * @return {string[][]}
 */
var groupAnagrams = function (strs) {
  const map = new Map();
  for (let s of strs) {
    const counts = Array.from(26).fill(0);
    for (let char of s) {
      const index = char.charCodeAt(0) - 97;
      counts[index]++;
    }
    const key = counts.join("#");
    if (!map.get(key)) {
      map.set(key, []);
    }
    map.get(key).push(s);
  }
  return Array.from(map.values());
};

console.log(groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]));
console.log(groupAnagrams([""]));
console.log(groupAnagrams(["a"]));
