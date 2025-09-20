package main

import "fmt"

func findTwoLargest(nums []int) []int {
	largest := nums[0]
	secondLargest := nums[1]
	// Swap values
	if secondLargest > largest {
		secondLargest, largest = largest, secondLargest
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] >= largest {
			secondLargest = largest
			largest = nums[i]
		} else if nums[i] > secondLargest && nums[i] < largest {
			secondLargest = nums[i]
		}
	}
	// Returns a slice of the largest and secondLargest values
	return []int{largest, secondLargest}
}

func main() {
	testCase1 := []int{15, 61, 25, 37, 52, 10, 47, 32, 73, 41, 88, 12, 64, 56, 29, 83, 98, 70, 90,
21}

	fmt.Println(findTwoLargest(testCase1))
}