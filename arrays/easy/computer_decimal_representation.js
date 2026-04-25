// You are given a positive integer n.

// A positive integer is a base-10 component if it is the product of a single digit from 1 to 9 and a non-negative power of 10. For example, 500, 30, and 7 are base-10 components, while 537, 102, and 11 are not.

// Express n as a sum of only base-10 components, using the fewest base-10 components possible.

// Return an array containing these base-10 components in descending order.

// Example 1:

// Input: n = 537

// Output: [500,30,7]

// Explanation:

// We can express 537 as 500 + 30 + 7. It is impossible to express 537 as a sum using fewer than 3 base-10 components.

// Example 2:

// Input: n = 102

// Output: [100,2]

// Explanation:

// We can express 102 as 100 + 2. 102 is not a base-10 component, which means 2 base-10 components are needed.

// Example 3:

// Input: n = 6

// Output: [6]

// Explanation:

// 6 is a base-10 component.

// Constraints:

//     1 <= n <= 109

/**
 * @param {number} n
 * @return {number[]}
 */
var decimalRepresentation = function (n) {
    const components = [];
    while (n > 0) {
        let powerOf10 = 1;
        // Find the largest power of 10 that is less than or equal to n
        // This ensures we are picking the most significant "digit" component first
        while (powerOf10 * 10 <= n) {
            powerOf10 *= 10;
        }

        // Determine the digit for this power of 10
        // For example, if n = 537 and powerOf10 = 100, digit = floor(537/100) = 5
        const digit = Math.floor(n / powerOf10);
        const component = digit * powerOf10; // This forms the base-10 component, e.g., 5 * 100 = 500

        components.push(component);
        n -= component; // Subtract the component to find the remainder
    }
    return components;
};
