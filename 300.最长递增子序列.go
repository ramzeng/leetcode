/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	var answer int

	n := len(nums)
	answers := make([]int, n)

	for i := 0; i < n; i++ {
		answers[i] = 1

		for j := 0; j < n; j++ {
			if nums[i] > nums[j] {
				answers[i] = max(answers[i], answers[j]+1)
			}
		}
	}

	for i, _ := range answers {
		answer = max(answer, answers[i])
	}

	return answer
}
// @lc code=end

