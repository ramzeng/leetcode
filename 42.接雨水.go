/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	var preMax, sufMax, area int
	n := len(height)
	left, right := 0, n-1

	for left < right {
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])

		if preMax < sufMax {
			area += preMax-height[left]
			left++
		} else {
			area += sufMax-height[right]
			right--
		}
	}

	return area
}
// @lc code=end

