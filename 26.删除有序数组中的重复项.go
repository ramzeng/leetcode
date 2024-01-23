/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
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

	return f(1)
}
// @lc code=end

