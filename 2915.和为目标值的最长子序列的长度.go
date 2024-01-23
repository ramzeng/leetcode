/*
 * @lc app=leetcode.cn id=2915 lang=golang
 *
 * [2915] 和为目标值的最长子序列的长度
 */

// @lc code=start
func lengthOfLongestSubsequence(nums []int, target int) int {
	answers := make([]int, target+1)

	for i, _ := range answers {
		answers[i] = math.MinInt
	}

	answers[0] = 0

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			answers[j] = max(answers[j], answers[j-nums[i]]+1)
		}
	}

	if answers[target] > 0 {
		return answers[target]
	} else {
		return -1
	}
}

// @lc code=end

