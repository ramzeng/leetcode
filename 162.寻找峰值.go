/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
func findPeakElement(nums []int) int {
	n := len(nums)

	left, right := 0, n-2

	for left <= right {
		mid := left+(right-left)/2

		if nums[mid] >= nums[mid+1] {
			right = mid-1
		} else {
			left = mid+1
		}
	}

	return left
}
// @lc code=end

