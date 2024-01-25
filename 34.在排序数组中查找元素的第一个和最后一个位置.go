/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	// 二分查找
	// left 是第一个大于等于 target 的下标
	// right 是第一个小于 target 的下标
	// right = left-1
	start := binarySearch(nums, target)

	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}

	end := binarySearch(nums, target+1) - 1

	return []int{start, end}
}

func binarySearch(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

// @lc code=end

