/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	// 找到最长的字符串作为基准值
	var prefix string
	for _, str := range strs {
		if len(str) > len(prefix) {
			prefix = str
		}
	}

	for _, str := range strs {
		var i int
		for ; i < len(str) && i < len(prefix); i++ {
			if str[i] != prefix[i] {
				break
			}
		}

		prefix = prefix[:i]
	}

	return prefix
}

// @lc code=end

