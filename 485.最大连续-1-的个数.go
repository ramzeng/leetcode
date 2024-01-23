/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续 1 的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	var maxCount, count int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			maxCount = max(maxCount, count)
		} else {
			count = 0
		}
	}

	return maxCount
}
// @lc code=end

