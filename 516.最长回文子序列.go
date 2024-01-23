/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 */

// @lc code=start
func longestPalindromeSubseq(s string) int {
	var q []byte

	for i := len(s) - 1; i >= 0; i-- {
		q = append(q, s[i])
	}

	return longestCommonSubsequence(s, string(q))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	answers := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		answers[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 不选 text1，选 text2
			// 选 text1，不选 text2
			answers[i][j] = max(answers[i-1][j], answers[i][j-1])

			// 判断是否一致
			if text1[i-1] == text2[j-1] {
				//
				answers[i][j] = max(answers[i][j], answers[i-1][j-1]+1)
			}
		}
	}

	return answers[m][n]
}

// @lc code=end

