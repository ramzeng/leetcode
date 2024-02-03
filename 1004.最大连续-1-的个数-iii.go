/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] 最大连续1的个数 III
 */

// @lc code=start
func longestOnes(nums []int, k int) int {
	count := 0
	slow, fast := 0, 0
	n := len(nums)
	answer := 0

	for ; fast < n; fast++ {
		if nums[fast] == 0 {
			count++
		}

		for count > k {
			if nums[slow] == 0 {
				count--
			}
			slow++
		}

		answer = max(answer, fast-slow+1)
	}

	return answer
}

// @lc code=end

