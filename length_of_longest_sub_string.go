package main

// 无重复字符的最长子串
// 左右指针，滑动窗口
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/submissions
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	m := map[byte]int{}

	right, result := 0, 0

	for i := 0; i < len(s); i++ {
		if i > 0 {
			// 左指针移动
			delete(m, s[i-1])
		}

		// 右指针移动
		for right < len(s) && m[s[right]] == 0 {
			m[s[right]]++
			right++
		}

		// 如果有重复值，保存结果，移动左指针
		result = max(right-i, result)
	}

	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
