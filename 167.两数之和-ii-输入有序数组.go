/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	left, right := 0, n-1

	for left < right {
		sum := numbers[left]+numbers[right]

		if sum == target {
			return []int{left+1, right+1}
		}

		if sum > target {
			right--
		} else {
			left++
		}
	}

	return nil
}
// @lc code=end

