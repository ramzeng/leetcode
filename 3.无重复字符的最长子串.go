/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	var slow, fast int
	var maxLength int

	visited := [128]int{}
	n := len(s)

	for ; fast < n; fast++ {
		visited[s[fast]]++
		for visited[s[fast]] > 1 {
			visited[s[slow]]--
			slow++
		}
		maxLength = max(maxLength, fast-slow+1)
	}

	return maxLength
}

// @lc code=end

