// Given the head of a singly linked list, return true if it is a or false otherwise.

// Example 1:

// Input: head = [1,2,2,1]
// Output: true

// Example 2:

// Input: head = [1,2]
// Output: false

// Constraints:

//     The number of nodes in the list is in the range [1, 105].
//     0 <= Node.val <= 9

// Follow up: Could you do it in O(n) time and O(1) space?

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {boolean}
 */
var isPalindrome = function (head) {
  if (!head || !head.next) return true;

  let slow = head;
  let fast = head;
  while (fast && fast.next) {
    slow = slow.next;
    fast = fast.next.next;
  }

  let prev = null;
  let curr = slow;

  while (curr) {
    let nextTemp = curr.next;
    curr.next = prev;
    prev = curr;
    curr = nextTemp;
  }

  let firstHalf = head;
  let secondHalf = prev;
  while (secondHalf) {
    if (firstHalf.val !== secondHalf.val) return false;
    firstHalf = firstHalf.next;
    secondHalf = secondHalf.next;
  }
  return true;
};

console.log(isPalindrome([1, 2, 2, 1]));
console.log(isPalindrome([1, 2]));
