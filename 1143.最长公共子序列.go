/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	answer := 0
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])

			if text1[i] == text2[j] {
				dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j]+1)
			}

			answer = max(answer, dp[i+1][j+1])
		}
	}

	return answer
}

// @lc code=end

