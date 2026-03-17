// In MATLAB, there is a handy function called reshape which can reshape an m x n matrix into a new one with a different size r x c keeping its original data.

// You are given an m x n matrix mat and two integers r and c representing the number of rows and the number of columns of the wanted reshaped matrix.

// The reshaped matrix should be filled with all the elements of the original matrix in the same row-traversing order as they were.

// If the reshape operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

// Example 1:

// Input: mat = [[1,2],[3,4]], r = 1, c = 4
// Output: [[1,2,3,4]]

// Example 2:

// Input: mat = [[1,2],[3,4]], r = 2, c = 4
// Output: [[1,2],[3,4]]

// Constraints:

//     m == mat.length
//     n == mat[i].length
//     1 <= m, n <= 100
//     -1000 <= mat[i][j] <= 1000
//     1 <= r, c <= 300

/**
 * @param {number[][]} mat
 * @param {number} r
 * @param {number} c
 * @return {number[][]}
 */
var matrixReshape = function (mat, r, c) {
  const m = mat.length;
  const n = mat[0].length;
  if (m * n !== r * c) return mat;

  const result = Array.from({ length: r }, () => new Array(c));
  let idx = 0;
  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      const newRow = Math.floor(idx / c);
      const newCol = idx % c;
      result[newRow][newCol] = mat[i][j];
      idx++;
    }
  }
  return result;
};

console.log(
  matrixReshape(
    [
      [1, 2],
      [3, 4]
    ],
    1,
    4
  )
);
console.log(
  matrixReshape(
    [
      [1, 2],
      [3, 4]
    ],
    2,
    4
  )
);
