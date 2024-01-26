/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	// 动态规划
	// 最长公共子序列 LCS
	m, n := len(text1), len(text2)

	// dp[i][j] text1 前 i 个字符，text2 前 j 个字符
	// 值是最长公共子序列长度

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])

			if text1[i-1] == text2[j-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}

	return dp[m][n]
}

// @lc code=end

