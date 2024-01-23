/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	f := func(nums []int, start, end int) int {
		n := len(nums)
		answers := make([]int, n+2)

		for i := start; i < end; i++ {
			answers[i+2] = max(answers[i+1], answers[i]+nums[i])
		}

		return answers[end+1]
	}

	n := len(nums)

	return max(nums[0]+f(nums, 2, n-1), f(nums, 1, n))
}


// @lc code=end

