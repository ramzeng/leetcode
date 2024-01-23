/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	var f func(k int) int

	f = func(k int) int {
		u := 0

		for i := 0; i < len(nums); i++ {
			if u < k || nums[u-k] != nums[i] {
				nums[u] = nums[i]
				u++
			}
		}

		return u
	}

	return f(2)
}
// @lc code=end

