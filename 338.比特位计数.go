/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(n int) []int {
	counter := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		count := 0
		for x := i; x > 0; x -= lowbit(x) {
			count++
		}
		counter[i] = count
	}

	return counter
}

func lowbit(x int) int {
	return x & (-x)
}

// @lc code=end

