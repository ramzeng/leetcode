/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	for _, num := range nums {
		target += num
	}

	if target < 0 || target%2 != 0 {
		return 0
	}

	target /= 2

	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, target+1)
	}

	dp[0][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < target+1; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= nums[i] {
				dp[i+1][j] = dp[i][j] + dp[i][j-nums[i]]
			}
		}
	}

	return dp[n][target]
}

// @lc code=end

