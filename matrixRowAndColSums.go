package main

import (
	"fmt"
)

func rowColSums(matrix [][]int) ([][]int, error) {
	// Edge Case: If matrix is nil, empty, or first row is empty
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil, fmt.Errorf("Matrix is nil, empty, or the first row doesn't exist")
	}
	// Initialize variables for processing
	rows := len(matrix)
	cols := len(matrix[0])
	rowSums := make([]int, rows)
	colSums := make([]int, cols)
	// Iterate the matrix, summing the rows and cols
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			rowSums[row] += matrix[row][col]
			colSums[col] += matrix[row][col]
		}
	}
	return [][]int{rowSums, colSums}, nil
}

func main() {
	matrixTest := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	sum, err := rowColSums(matrixTest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("The sum of the array is [Row Sums: %v, Column Sums: %v]\n", sum[0], sum[1])

}
