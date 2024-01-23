/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	n := len(height)
	maxArea := math.MinInt
	left, right := 0, n-1

	for left <= right {
		maxArea = max(maxArea, min(height[left], height[right])*(right-left))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
// @lc code=end

