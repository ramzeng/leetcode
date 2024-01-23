/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 0; i < len(strs); i++ {
		var j int

		for ; j < len(prefix) && j < len(strs[i]); j++ {
			if strs[i][j] != prefix[j] {
				break
			}
		}

		prefix = prefix[:j]
	}

	return prefix
}

// @lc code=end

