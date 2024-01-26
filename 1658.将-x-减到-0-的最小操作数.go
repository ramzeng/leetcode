/*
 * @lc app=leetcode.cn id=1658 lang=golang
 *
 * [1658] 将 x 减到 0 的最小操作数
 */

// @lc code=start
func minOperations(nums []int, x int) int {
	target := -x

	for _, num := range nums {
		target += num
	}

	if target < 0 {
		return -1
	}

	var slow, fast int
	var answer, sum int
	answer = -1

	n := len(nums)

	for ; fast < n; fast++ {
		sum += nums[fast]

		for sum > target {
			sum -= nums[slow]
			slow++
		}

		if sum == target {
			answer = max(answer, fast-slow+1)
		}
	}

	if answer == -1 {
		return -1
	}

	return n - answer
}

// @lc code=end

