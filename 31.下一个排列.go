/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	// 数组，模板题
	// 从倒数第二个开始找，第一个破坏升序的数的下标
	// 从右往左找第一个大于该数的数，进行互换
	// 再把 i+1 到 n-1 这一段数组翻转

	n := len(nums)
	i := n - 2

	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		for j := n - 1; j >= 0; j-- {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}

	left, right := i+1, n-1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// @lc code=end

