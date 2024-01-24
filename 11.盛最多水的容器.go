/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	// 相向指针

	n := len(height)
	maxArea := math.MinInt
	left, right := 0, n-1

	for left < right {
		area := min(height[left], height[right]) * (right-left)
		maxArea = max(maxArea, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
// @lc code=end

