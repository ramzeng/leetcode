package main

import (
	"sort"
)

// 三数之和
// 排序 + 双指针法
// https://leetcode.cn/problems/3sum/submissions/
func threeSum(nums []int) [][]int {
	var results [][]int

	if len(nums) < 3 {
		return results
	}

	// 排序
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		// 如果最左侧的值都大于 0，直接结束循环
		if nums[i] > 0 {
			break
		}

		// 注意 i 溢出问题，所以要判断 i > 0
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 左右指针
		left := i + 1
		right := len(nums) - 1

		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				results = append(results, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// 这里不能 continue，需要继续下面的逻辑，否则会造成死循环
			}

			// 大于 0 右指针向左移动，小于 0 左指针向右移动
			if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return results
}
