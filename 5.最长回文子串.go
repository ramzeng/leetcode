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

	// 中心扩散法，枚举中点
	// left, right 记录答案区间
	var left, right int

	for i := 0; i < len(s); i++ {
		// 奇数情况
		x := expandFromCenter(s, i, i)
		// 偶数情况
		y := expandFromCenter(s, i, i+1)
		// 求最大值
		z := max(x, y)

		if z > right-left {
			// 需要写 z-1，不然奇数会有问题
			left = i - (z-1)/2
			right = i + z/2
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

	// 去掉中间字符
	return right - left - 1
}

// @lc code=end

