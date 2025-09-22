package main

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func InOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	left := InOrderTraversal(root.Left)
	right := InOrderTraversal(root.Right)

	result = append(result, left...)
	result = append(result, root.Value)
	result = append(result, right...)

	return result
}

func flattenBST(root *TreeNode, min, max int) []int {
	if root == nil {
		return []int{}
	}
	results := InOrderTraversal(root)
	var newResult []int
	for _, num := range results {
		if num >= min && num <= max {
			newResult = append(newResult, num)
		}
	}
	return newResult
}

func main() {
	root := &TreeNode{
		Value: 100,
		Left: &TreeNode{
			Value: 50,
			Left: &TreeNode{
				Value: 25,
				Left: &TreeNode{
					Value: 10,
				},
				Right: &TreeNode{
					Value: 30,
				},
			},
			Right: &TreeNode{
				Value: 75,
				Left: &TreeNode{
					Value: 60,
				},
				Right: &TreeNode{
					Value: 80,
				},
			},
		},
		Right: &TreeNode{
			Value: 150,
			Left: &TreeNode{
				Value: 125,
				Left: &TreeNode{
					Value: 110,
				},
				Right: &TreeNode{
					Value: 130,
				},
			},
			Right: &TreeNode{
				Value: 175,
				Left: &TreeNode{
					Value: 160,
				},
				Right: &TreeNode{
					Value: 190,
				},
			},
		},
	}
	min := 60
	max := 150
	fmt.Println(flattenBST(root, min, max))
}
