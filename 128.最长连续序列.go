/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)

	for _, num := range nums {
		m[num] = true
	}

	maxCount := 0

	for num, _ := range m {
		if m[num-1] {
			continue
		}

		current := num
		count := 1

		for m[current+1] {
			current++
			count++
		}

		maxCount = max(maxCount, count)
	}

	return maxCount
}

// @lc code=end

