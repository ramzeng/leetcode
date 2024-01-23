/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 */

// @lc code=start
func triangleNumber(nums []int) int {
	sort.Ints(nums)

	n := len(nums)
	answer := 0

	for i := 2; i < n; i++ {
		left, right := 0, i-1

		for left < right {
			if nums[left]+nums[right] > nums[i] {
				answer += right-left
				right--
			} else {
				left++
			}
		}
	}

	return answer
}
// @lc code=end

