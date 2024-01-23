/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	var slow, fast int

	for ; fast < len(nums); fast++ {
		if nums[fast] == val {
			continue
		}

		nums[slow] = nums[fast]
		slow++
	}

	return slow
}
// @lc code=end

