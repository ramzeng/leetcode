/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-2

	for left <= right {
		mid := left + (right-left)/2

		// 跟右侧第一个比
		if nums[mid] > nums[right+1] {
			left = mid + 1
		} else if nums[mid] < nums[right+1] {
			right = mid - 1
		} else {
			right--
		}
	}

	return nums[left]
}

// @lc code=end

