// Design an algorithm to compute the sum of each row and each column
// Returns two arrays, one for row sums and one for column sums
// Algorithm sums all the rows from top to bottom. Then it sums all the cols from left to right.

const rowColSums = (matrix) => {
  // Edge Case: If not a matrix data structure, empty, or first row length is empty
  if (!matrix || matrix.length === 0 || matrix[0].length === 0) {
    return [[], []];
  }
  // Initialize variables(rows and cols) to set row and length for proper algo loop sizing
  let rows = matrix.length;
  let cols = matrix[0].length;
  // Initialize variables(rowSums and colSums to set contiguous blocks of memory for storing sums) and fill with 0's as placeholder initially
  let rowSums = new Array(rows).fill(0);
  let colSums = new Array(cols).fill(0);
  // Iterate the matrix. Each pass sum the first row/col and repeat until matrix end is reached
  for (let row = 0; row < rows; row++) {
    for (let col = 0; col < cols; col++) {
      rowSums[row] += matrix[row][col];
      colSums[col] += matrix[row][col];
    }
  }
  return [rowSums, colSums];
};

// Testing

let matrixTest = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9],
];
console.log(rowColSums(matrixTest));
