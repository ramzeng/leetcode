/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	count, preSum := 0, 0
	hash := make(map[int]int)
	hash[0] = 1
	n := len(nums)

	for i := 0; i < n; i++ {
		preSum += nums[i]

		if c, ok := hash[preSum-k]; ok {
			count += c
		}

		hash[preSum]++
	}

	return count
}

// @lc code=end

