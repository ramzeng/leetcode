/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, amount+1)
	}

	dp[0][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < amount+1; j++ {
			if j >= coins[i] {
				dp[i+1][j] = dp[i][j] + dp[i+1][j-coins[i]]
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}

	return dp[n][amount]
}

// @lc code=end

