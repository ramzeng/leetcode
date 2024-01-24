/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	n := len(nums)
	// 最后一个值作为基准值，不能动，所以从倒数第二个开始
	left, right := 0, n-2

	for left <= right {
		mid := left + (right-left)/2

		// 跟最后一个数比
		if nums[mid] >= nums[n-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return nums[left]
}

// @lc code=end

