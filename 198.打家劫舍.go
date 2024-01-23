/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	answers := make([]int, n+2)

	for i := 0; i < n; i++ {
		answers[i+2] = max(answers[i+1], answers[i]+nums[i])
	}

	return answers[n+1]
}
// @lc code=end

