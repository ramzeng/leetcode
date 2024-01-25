/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	// 相向双指针，维护前缀最大和后缀最大
	var preMax, sufMax, answer int

	left, right := 0, len(height)-1

	for left <= right {
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])

		if preMax < sufMax {
			answer += preMax - height[left]
			left++
		} else {
			answer += sufMax - height[right]
			right--
		}
	}

	return answer
}

// @lc code=end

