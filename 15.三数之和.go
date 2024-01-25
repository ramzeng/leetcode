package main

import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	// 数组，求 target，组合
	// 排序
	sort.Ints(nums)

	// 相向双指针
	var paths [][]int
	n := len(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				paths = append(paths, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for right > left && nums[right] == nums[right-1] {
					right--
				}
			}

			if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return paths
}

// @lc code=end
