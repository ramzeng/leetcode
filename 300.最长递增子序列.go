/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	// 动态规划
	// 最长递增子序列 LIS
	// dp[i]，前 i 个数字，最长递增子序列
	n := len(nums)
	answer := 0
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	for i := 0; i < n; i++ {
		answer = max(answer, dp[i])
	}

	return answer
}

// @lc code=end

