/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	i := findMin(nums)
	n := len(nums)

	var left, right int

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
		return -1
	} else {
		return left
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

