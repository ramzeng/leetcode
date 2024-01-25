/*
 * @lc app=leetcode.cn id=2859 lang=golang
 *
 * [2859] 计算 K 置位下标对应元素的和
 */

// @lc code=start
func sumIndicesWithKSetBits(nums []int, k int) int {
	var sum int

	for i := 0; i < len(nums); i++ {
		count := k
		j := i
		for ; j != 0 && count >= -1; count-- {
			j -= lowbit(j)
		}

		if count == 0 {
			sum += nums[i]
		}
	}

	return sum
}

func lowbit(x int) int {
	return x & (-x)
}

// @lc code=end

