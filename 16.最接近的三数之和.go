package main

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var diff, minDiff, answer int

	minDiff = math.MaxInt
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

		for left < right {
			sum := nums[i]+nums[left]+nums[right]
			diff = abs(sum-target)
			if diff < minDiff {
				minDiff = diff
				answer = sum
			}

			if nums[i]+nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}
		}
	}

	return answer
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// @lc code=end

