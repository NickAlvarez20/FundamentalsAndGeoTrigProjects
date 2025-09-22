const findTargetWithinMatrix = (matrix, target) => {
  // Edge Case: If matrix is null or length is zero
  if (!matrix || matrix.length === 0 || isNaN(target)) {
    return "The Matrix has you Neo, enter a matrix or valid target number";
  }
  // Find a given target by iterating the matrix to see if it exists within
  for (let row = 0; row < matrix.length; row++) {
    for (let col = 0; col < matrix[row].length; col++) {
      if (matrix[row][col] === target) {
        return `Target found at ${[row, col]}`;
      }
    }
  }
  return `Target Out of Bounds or Not Found, ${[-1, -1]}`;
};

let matrixTest = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9],
];
console.log(findTargetWithinMatrix(matrixTest, 111));
