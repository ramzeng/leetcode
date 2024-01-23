/*
 * @lc app=leetcode.cn id=2824 lang=golang
 *
 * [2824] 统计和小于目标的下标对数目
 */

// @lc code=start
func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	answer := 0
	left, right := 0, n-1

	for left < right {
		if nums[left]+nums[right] < target {
			answer += right-left
			left++
		} else {
			right--
		}
	}

	return answer
}
// @lc code=end

