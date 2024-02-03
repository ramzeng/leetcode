/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	n := len(nums)
	start := sort.SearchInts(nums, target)

	if start == n || nums[start] != target {
		return []int{-1, -1}
	}

	end := sort.SearchInts(nums, target+1)

	return []int{start, end - 1}
}

// @lc code=end

