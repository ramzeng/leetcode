/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// 一维动态规划
	// 是指动态规划问题的状态只有一个维度
	// 只需要一个一维数组存储中间状态
	n := len(nums)
	dp := make([]int, n)

	// 表示前 i 个数，最大连续子数组和
	// dp[i] i 表示前 i 个数，值表示前 i 个数最大连续子数组和
	dp[0] = nums[0]
	maxSum := nums[0]

	for i := 1; i < n; i++ {
		// 维护中间状态
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		// 计算中间态的最大值
		maxSum = max(maxSum, dp[i])
	}

	return maxSum
}

// @lc code=end

