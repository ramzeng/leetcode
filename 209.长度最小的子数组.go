/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	var slow, fast, sum int
	answer := math.MaxInt

	n := len(nums)

	for ; fast < n; fast++ {
		sum += nums[fast]

		for sum >= target {
			answer = min(answer, fast-slow+1)
			sum -= nums[slow]
			slow++
		}
	}

	if answer < math.MaxInt {
		return answer
	} else {
		return 0
	}
}
// @lc code=end

