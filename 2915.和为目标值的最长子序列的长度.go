/*
 * @lc app=leetcode.cn id=2915 lang=golang
 *
 * [2915] 和为目标值的最长子序列的长度
 */

// @lc code=start
func lengthOfLongestSubsequence(nums []int, target int) int {
	// 0 / 1 背包，最大值
	n := len(nums)
	// dp[i][j] 在前 i 个选，和为 j 的情况下，数字的个数
	dp := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 0
	}

	for i := 0; i < target+1; i++ {
		dp[0][i] = math.MinInt
	}

	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j <= target; j++ {
			dp[i+1][j] = dp[i][j]

			if j >= nums[i] {
				dp[i+1][j] = max(dp[i][j], dp[i][j-nums[i]]+1)
			}
		}
	}

	if dp[n][target] > 0 {
		return dp[n][target]
	} else {
		return -1
	}
}

// @lc code=end

