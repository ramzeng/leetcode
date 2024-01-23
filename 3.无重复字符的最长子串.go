/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	var viewed [128]int
	var slow, fast, answer int

	for ; fast < len(s); fast++ {
		viewed[s[fast]]++

		for viewed[s[fast]] > 1 {
			viewed[s[slow]]--
			slow++
		}

		answer = max(answer, fast-slow+1)
	}

	return answer
}
// @lc code=end

