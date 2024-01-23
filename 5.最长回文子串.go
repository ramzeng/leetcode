/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var start, end int

	for i := 0; i < len(s); i++ {
		x := expandAroundCenter(s, i, i)
		y := expandAroundCenter(s, i, i+1)
		z := max(x, y)

		if z > end-start {
			start = i - (z-1)/2
			end = i + z/2
		}
	}

	return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for ; left >= 0 && right < len(s); left, right = left-1, right+1 {
		if s[left] != s[right] {
			break
		}
	}

	return right - left - 1
}

// @lc code=end

