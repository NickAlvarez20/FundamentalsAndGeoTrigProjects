package main

import (
	"fmt"
)

func findTargetWithinMatrix(matrix [][]int64, target int64) ([]int, error) {
	// Edge Case: If matrix data type is nil or empty
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil, fmt.Errorf("invalid data type or empty matrix")
	}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == target {
				return []int{row,col}, nil
			}

		}
	}
	// Target not found
	return nil, fmt.Errorf("Target not found within matrix")
}

func main() {
	// Create a test matrix
	matrixTest := [][]int64{{}, {4, 5, 6}, {7, 8, 9}}
	// Test findTargetWithinMatrix func
	indices, err := findTargetWithinMatrix(matrixTest, 9)
	if err != nil{
		fmt.Printf("Error:", err)
		return
	}
	fmt.Printf("Target found at indices: [%d, %d]\n", indices[0], indices[1])

}
