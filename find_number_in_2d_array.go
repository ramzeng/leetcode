package main

import (
	"fmt"
)

func main() {
	fmt.Println(findNumberIn2DArray([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 5))
}

// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}

	length := len(matrix[0])

	if length < 1 {
		return false
	}

	for _, nums := range matrix {
		if target < nums[0] || target > nums[length-1] {
			continue
		}

		for _, num := range nums {
			if target == num {
				return true
			}
		}
	}

	return false
}
