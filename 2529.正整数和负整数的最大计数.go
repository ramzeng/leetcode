/*
 * @lc app=leetcode.cn id=2529 lang=golang
 *
 * [2529] 正整数和负整数的最大计数
 */

// @lc code=start
func maximumCount(nums []int) int {
	var posCount, negCount int

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			negCount++
		}

		if nums[i] > 0 {
			posCount++
		}
	}

	return max(posCount, negCount)
}
// @lc code=end

