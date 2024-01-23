/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] 最大连续1的个数 III
 */

// @lc code=start
func longestOnes(nums []int, k int) int {
	var slow, fast int
	var count, maxCount int

	for ; fast < len(nums); fast++ {
		if nums[fast] == 0 {
			count++
		}

		for count > k {
			if nums[slow] == 0 {
				count--
			}

			slow++
		}

		maxCount = max(maxCount, fast-slow+1)
	}

	return maxCount
}
// @lc code=end

