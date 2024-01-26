/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	var slow, fast int
	var sum, answer int

	n := len(nums)
	answer = n + 1

	for ; fast < n; fast++ {
		sum += nums[fast]

		for sum >= target {
			answer = min(answer, fast-slow+1)
			sum -= nums[slow]
			slow++
		}
	}

	if answer > n {
		return 0
	} else {
		return answer
	}
}

// @lc code=end

