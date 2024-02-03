/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+2)

	for i := 0; i < n; i++ {
		dp[i+2] = max(dp[i+1], dp[i]+nums[i])
	}

	return dp[n+1]
}

// @lc code=end

