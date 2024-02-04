/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续 1 的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	maxCount := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] != 1 {
			maxCount = max(maxCount, count)
			count = 0
			continue
		}

		count++
	}

	maxCount = max(maxCount, count)

	return maxCount
}

// @lc code=end

