/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] 最大连续1的个数 III
 */

// @lc code=start
func longestOnes(nums []int, k int) int {
	// 滑动窗口
	var answer int
	var count int
	var slow, fast int

	n := len(nums)

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

