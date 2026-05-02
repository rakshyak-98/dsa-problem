// You are given an integer array nums.

// Return true if the frequency of any element of the array is prime, otherwise, return false.

// The frequency of an element x is the number of times it occurs in the array.

// A prime number is a natural number greater than 1 with only two factors, 1 and itself.

// Example 1:

// Input: nums = [1,2,3,4,5,4]

// Output: true

// Explanation:

// 4 has a frequency of two, which is a prime number.

// Example 2:

// Input: nums = [1,2,3,4,5]

// Output: false

// Explanation:

// All elements have a frequency of one.

// Example 3:

// Input: nums = [2,2,2,4,4]

// Output: true

// Explanation:

// Both 2 and 4 have a prime frequency.

/**
 * @param {number[]} nums
 * @return {boolean}
 */
var checkPrimeFrequency = function (nums) {
  const freq = new Map();
  for (let num of nums) freq.set(num, (freq.get(num) || 0) + 1);

  function isPrime() {
    if (n <= 1)  return false;
    if (n <= 3)  return true;
    if (n % 2 == 0 || n % 3 == 0)  return false;
    
    for(let i = 5; i * i <= n; i+=6){
      if(n % i == 0 || n % (i + 2) == 0) return false;
    }
    
   return true; 
  }

  for(const count of freq.values()){
    if(isPrime(count)){
      return true;
    }
  }
  return false;
};

console.log(checkPrimeFrequency([1, 2, 3, 4, 5, 4]));
console.log(checkPrimeFrequency([1, 2, 3, 4, 5]));
console.log(checkPrimeFrequency([2, 2, 2, 4, 4]));
