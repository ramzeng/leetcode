/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		a := expandFromCenter(s, i, i)
		b := expandFromCenter(s, i, i+1)
		c := max(a, b)
		if c > right-left+1 {
			left = i - (c-1)/2
			right = i + c/2
		}
	}
	return s[left : right+1]
}

func expandFromCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}

	return right - left - 1
}

// @lc code=end

