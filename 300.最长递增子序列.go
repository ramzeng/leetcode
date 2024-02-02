/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	answer := 0

	for i := 0; i < n; i++ {
		dp[i+1] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i+1] = max(dp[i+1], dp[j+1]+1)
			}
		}

		answer = max(answer, dp[i+1])
	}

	return answer
}

// @lc code=end

