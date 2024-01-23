/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			var j int

			for ; j < len(needle) && i+j < len(haystack); j++ {
				if haystack[i+j] != needle[j] {
					break
				}
			}

			if j == len(needle) {
				return i
			}
		}
	}

	return -1
}
// @lc code=end

