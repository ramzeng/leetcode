/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 */

// @lc code=start
func longestPalindromeSubseq(s string) int {
	n := len(s)
	left, right := 0, n-1
	p := []byte(s)

	for left < right {
		p[left], p[right] = p[right], p[left]
		left++
		right--
	}

	return longestSubseq(s, string(p))
}

func longestSubseq(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	answer := 0

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

