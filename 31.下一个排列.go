/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int)  {
	n := len(nums)

	// 从倒数第二个开始找，找到非升序排列的下标
	i := n-2

	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 如果找到了
	if i >= 0 {
		j := n-1

		for j >= 0 && nums[j] <= nums[i] {
			j--
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	left, right := i+1, n-1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
// @lc code=end

