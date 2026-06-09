// There are n persons on a social media website. You are given an integer array ages where ages[i] is the age of the ith person.

// A Person x will not send a friend request to a person y (x != y) if any of the following conditions is true:

//     age[y] <= 0.5 * age[x] + 7
//     age[y] > age[x]
//     age[y] > 100 && age[x] < 100

// Otherwise, x will send a friend request to y.

// Note that if x sends a request to y, y will not necessarily send a request to x. Also, a person will not send a friend request to themself.

// Return the total number of friend requests made.

// Example 1:

// Input: ages = [16,16]
// Output: 2
// Explanation: 2 people friend request each other.

// Example 2:

// Input: ages = [16,17,18]
// Output: 2
// Explanation: Friend requests are made 17 -> 16, 18 -> 17.

// Example 3:

// Input: ages = [20,30,100,110,120]
// Output: 3
// Explanation: Friend requests are made 110 -> 100, 120 -> 110, 120 -> 100.

// Constraints:

//     n == ages.length
//     1 <= n <= 2 * 104
//     1 <= ages[i] <= 120

/**
 * @param {number[]} ages
 * @return {number}
 */
var numFriendRequests = function (ages) {
  const count = new Array(121).fill(0);
  for (const age of ages) {
    count[age]++;
  }

  let totalRequests = 0;

  for (let ageX = 1; ageX <= 120; ageX++) {
    if (count[ageX] === 0) continue;

    for (let ageY = 1; ageY <= 120; ageY++) {
      if (count[ageY] == 0) continue;
      if (ageY <= 0.5 * ageX + 7 || ageY > ageX || (ageY > 100 && ageX < 100)) {
        continue;
      }

      if (ageX === ageY) {
        totalRequests += count[ageX] * (count[ageX] - 1);
      } else {
        totalRequests += count[ageX] * count[ageY];
      }
    }
  }
  return totalRequests;
};

console.log(numFriendRequests([16, 16]));
console.log(numFriendRequests([16, 17, 18]));
console.log(numFriendRequests([20,30,100,110,120]));
