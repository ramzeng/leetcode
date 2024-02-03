/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	if n == 1 {
		return n
	}

	dp := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		dp[i] = math.MaxInt
	}

	dp[0] = 0

	for i := 1; i < n+1; i++ {
		for j := i * i; j < n+1; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}

	return dp[n]
}

// @lc code=end

