/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	k := 1
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

