/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search(nums []int, target int) bool {
	i := findMin(nums)
	n := len(nums)
	left, right := 0, 0

	if target > nums[n-1] {
		left, right = 0, i-1
	} else {
		left, right = i, n-1
	}

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left == n || nums[left] != target {
		return false
	} else {
		return true
	}
}

func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-2
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right+1] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return left
}

// @lc code=end

