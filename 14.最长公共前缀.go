/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	longestPrefix := slices.Max(strs)

	for _, str := range strs {
		i, j := 0, 0
		for i < len(longestPrefix) && j < len(str) {
			if longestPrefix[i] != str[j] {
				break
			}

			i++
			j++
		}

		longestPrefix = longestPrefix[:i]
	}

	return longestPrefix
}

// @lc code=end

