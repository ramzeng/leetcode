package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}

func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}

	left, right := 0, len(nums)-1

	var answer []int

	for left < right {
		middle := (left + right) >> 1

		if nums[middle] >= target {
			right = middle
		} else {
			left = middle + 1
		}
	}

	if nums[left] != target {
		return []int{-1, -1}
	} else {
		answer = append(answer, left)

		left, right = 0, len(nums)-1

		for left < right {
			middle := (left + right + 1) >> 1

			if nums[middle] <= target {
				left = middle
			} else {
				right = middle - 1
			}
		}

		answer = append(answer, left)
	}

	return answer
}
