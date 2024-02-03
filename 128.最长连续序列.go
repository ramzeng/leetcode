/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	store := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		store[nums[i]] = true
	}

	var maxLength, length int

	for num, _ := range store {
		// 剪枝，如果上一个数算过了，这个数字不用算了
		if store[num-1] {
			continue
		}

		length = 1

		for store[num+1] {
			length++
			num++
		}

		maxLength = max(maxLength, length)
	}

	return maxLength
}

// @lc code=end

