/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	// 数组，原地哈希
	n := len(nums)

	// 当前数在 1 到 n 之间，且不在正确的位置上，且目标位置上的数不等于当前数
	// 则将当前数交换到正确的位置上
	for i := 0; i < n; i++ {
		// nums[i]-1 相当于是个 key
		for nums[i] > 0 && nums[i] < n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// 找到第一个不在正确位置上的数，返回其应在的位置
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

// @lc code=end

