/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	n := len(nums)
	min := findMin(nums)
	left, right := 0, 0

	if nums[n-1] >= target {
		left, right = min, n-1
	} else {
		left, right = 0, min-1
	}

	index := sort.SearchInts(nums[left:right+1], target)

	index += left

	if index == len(nums) || nums[index] != target {
		return -1
	} else {
		return index
	}
}

func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-2
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= nums[n-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// @lc code=end

