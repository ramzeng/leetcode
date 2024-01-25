/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	// 字符串，字符串查找
	n := len(haystack)
	for i := 0; i < n; i++ {
		// 判断第一个字符是否一致
		if haystack[i] != needle[0] {
			continue
		}

		// 判断剩余长度够不够匹配
		if n-i < len(needle) {
			return -1
		}

		// 往后匹配
		j := 0
		for ; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}

		// 是否完全匹配
		if j == len(needle) {
			return i
		}
	}

	return -1
}

// @lc code=end

