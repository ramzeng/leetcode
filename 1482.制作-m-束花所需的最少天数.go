/*
 * @lc app=leetcode.cn id=1482 lang=golang
 *
 * [1482] 制作 m 束花所需的最少天数
 */

// @lc code=start
func minDays(bloomDay []int, m int, k int) int {
	left, right := 1, slices.Max(bloomDay)

	for left <= right {
		mid := left + (right-left)/2

		if canMake(bloomDay, m, k, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if canMake(bloomDay, m, k, left) {
		return left
	}

	return -1
}

// m 多少束花
// k 连续多少束
func canMake(bloomDay []int, m int, k int, days int) bool {
	// p 可以完成多少束
	// q 已经连续多少朵
	p, q := 0, 0
	for _, day := range bloomDay {
		if day <= days {
			q++
			// 连续 k 朵
			if q == k {
				// 完成 1 束
				p++
				// 重置
				q = 0
			}
		} else {
			// 断了，重置
			q = 0
		}
	}

	return p >= m
}

// @lc code=end

