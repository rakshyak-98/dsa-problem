// Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.

// Note that elements beyond the length of the original array are not written. Do the above modifications to the input array in place and do not return anything.

 

// Example 1:

// Input: arr = [1,0,2,3,0,4,5,0]
// Output: [1,0,0,2,3,0,0,4]
// Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

// Example 2:

// Input: arr = [1,2,3]
// Output: [1,2,3]
// Explanation: After calling your function, the input array is modified to: [1,2,3]

 

// Constraints:

    1 <= arr.length <= 104
    0 <= arr[i] <= 9
/**
 * @param {number[]} arr
 * @return {void} Do not return anything, modify arr in-place instead.
 */
var duplicateZeros = function (arr) {
  let possible_dups = 0;
  let length_ = arr.length - 1;

  for (let left = 0; left <= length_ - possible_dups; left++) {
    if (arr[left] == 0) {
      if (left === length_ - possible_dups) {
        arr[length_] = 0;
       length_ -= 1;
        break;
      }
      possible_dups++;
    }
  }

  let last = length_ - possible_dups;
  for (let i = last; i >= 0; i--) {
    if (arr[i] === 0) {
      arr[i + possible_dups] = 0;
      possible_dups--;
      arr[i + possible_dups] = 0;
    } else {
      arr[i + possible_dups] = arr[i];
    }
  }
};

var arr = [1, 0, 2, 3, 0, 4, 5, 0];
duplicateZeros(arr);
console.log(arr);

var arr = [1, 2, 3];
duplicateZeros(arr);
console.log(arr);
