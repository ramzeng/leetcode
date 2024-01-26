/*
 * @lc app=leetcode.cn id=713 lang=golang
 *
 * [713] 乘积小于 K 的子数组
 */

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	var slow, fast int
	var prod int
	var answer int

	prod = 1
	n := len(nums)

	for ; fast < n; fast++ {
		prod *= nums[fast]

		for prod >= k {
			prod /= nums[slow]
			slow++
		}

		answer += fast - slow + 1
	}

	return answer
}

// @lc code=end

