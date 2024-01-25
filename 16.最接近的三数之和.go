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
	// 数组，求 target，最小 diff
	// 排序
	sort.Ints(nums)

	// 相向双指针
	minDiff, diff, n, answer := math.MaxInt, 0, len(nums), 0

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			diff = abs(sum - target)

			if diff < minDiff {
				minDiff = diff
				answer = sum
			}

			if sum > target {
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
