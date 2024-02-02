/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// 前缀和解法
	// answer := math.MinInt
	// minPreSum := 0
	// currentPreSum := 0

	// for i := 0; i < len(nums); i++ {
	// 	currentPreSum += nums[i]
	// 	answer = max(answer, currentPreSum-minPreSum)
	// 	minPreSum = min(minPreSum, currentPreSum)
	// }

	// return answer

	// 动态规划
	n := len(nums)
	dp := make([]int, n+1)
	answer := math.MinInt

	for i := 0; i < n; i++ {
		dp[i+1] = nums[i]

		if dp[i] > 0 {
			dp[i+1] = dp[i] + nums[i]
		}

		answer = max(answer, dp[i+1])
	}

	return answer

	// 动态规划，空间优化
	// n := len(nums)
	// current := 0
	// answer := math.MinInt

	// for i := 0; i < n; i++ {
	// 	current = max(current+nums[i], nums[i])
	// 	answer = max(answer, current)
	// }

	// return answer
}

// @lc code=end

