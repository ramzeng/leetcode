/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, amount+1)
	}

	for i := 0; i < amount+1; i++ {
		dp[0][i] = amount+1
	}

	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j < amount+1; j++ {
			if j >= coins[i] {
				dp[i+1][j] = min(dp[i][j], dp[i+1][j-coins[i]]+1)
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}

	if dp[n][amount] == amount+1 {
		return -1
	} else {
		return dp[n][amount]
	}
}

// @lc code=end

