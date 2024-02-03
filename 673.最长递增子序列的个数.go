/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] 最长递增子序列的个数
 */

// @lc code=start
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	counter := make([]int, n)
	maxLength := 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		counter[i] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] == dp[j]+1 {
					counter[i] += counter[j]
				} else if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					counter[i] = counter[j]
				}
			}
		}

		maxLength = max(maxLength, dp[i])
	}

	count := 0

	for i := 0; i < n; i++ {
		if dp[i] == maxLength {
			count += counter[i]
		}
	}

	return count
}

// @lc code=end

