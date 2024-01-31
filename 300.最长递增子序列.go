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
	// 存储中间状态
	dp := make([]int, n)

	// 从第一个数字开始遍历
	for i := 0; i < n; i++ {
		dp[i] = 1

		// 遍历 i 之前的数字
		for j := 0; j < i; j++ {
			// 如果 i 之前的数字比 i 小，那么 dp[i] = dp[j] + 1
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// 遍历 dp，找到最大值
	for i := 0; i < n; i++ {
		answer = max(answer, dp[i])
	}

	return answer
}

// @lc code=end

