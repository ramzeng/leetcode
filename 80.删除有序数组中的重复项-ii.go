/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	k := 2
	n := len(nums)
	slow, fast := 0, 0

	for ; fast < n; fast++ {
		if slow < k || nums[slow-k] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

// @lc code=end

