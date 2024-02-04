/*
 * @lc app=leetcode.cn id=1482 lang=golang
 *
 * [1482] 制作 m 束花所需的最少天数
 */

// @lc code=start
func minDays(bloomDay []int, m int, k int) int {
	// m 束花，需要相邻的 k 朵
	right := slices.Max(bloomDay)
	left := 0

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
	} else {
		return -1
	}
}

func canMake(bloomDay []int, m, k, days int) bool {
	count := 0
	flowersCount := 0

	for _, day := range bloomDay {
		if day <= days {
			flowersCount++

			if flowersCount == k {
				count++
				flowersCount = 0
			}
		} else {
			flowersCount = 0
		}
	}

	return count >= m
}

// @lc code=end

