/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	// 1,2,3 -> 1,3,2
	n := len(nums)

	i := n - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}

	if i >= 0 {
		j := n - 1
		for ; j >= 0; j-- {
			if nums[j] > nums[i] {
				break
			}
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

