package main

import(
	"fmt"
	"math"
)

func oddSumMaxPair(nums[] int) []int{
	var result []int
	var currMaxPair, pairSum int
	for i:= 0; i<len(nums);i++{
		for j:= i; j<len(nums)-1; j++{
			pairSum = nums[i] + nums[j]
			if math.Abs(float64(pairSum%2))==1 && (pairSum > currMaxPair || len(result) == 0){
				currMaxPair = pairSum
				result = []int{nums[i], nums[j]}
			}
		}
	}
	return result
}

func main(){
	testCase1 := []int{-3, -2, -1, -4, -5}
	fmt.Println(oddSumMaxPair(testCase1))
}