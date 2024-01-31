/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// 原地哈希
		// 解决找到数组中出现两个
		// 缺失的第一个正数等等
		for nums[i] > 0 && nums[i] < n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	var answer []int

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			answer = append(answer, nums[i])
		}
	}

	return answer
}

// @lc code=end

