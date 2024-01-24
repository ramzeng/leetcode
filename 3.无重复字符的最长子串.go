/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// 同向快慢指针
	// viewed 数组记录之前的字符串个数
	// slow，fast 快慢指针
	// key 为 ascii，value 为个数
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

