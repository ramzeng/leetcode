/*
 * @lc app=leetcode.cn id=1658 lang=golang
 *
 * [1658] 将 x 减到 0 的最小操作数
 */

// @lc code=start
func minOperations(nums []int, x int) int {
	target := -x

	for i := 0; i < len(nums); i++ {
		target += nums[i]
	}

	if target < 0 {
		return -1
	}

	var slow, fast, sum int
	maxCount := -1

	for ; fast < len(nums); fast++ {
		sum += nums[fast]

		for sum > target {
			sum -= nums[slow]
			slow++
		}

		if sum == target {
			maxCount = max(maxCount, fast-slow+1)
		}
	}

	if maxCount < 0 {
		return -1
	}

	return len(nums)-maxCount
}
// @lc code=end

