/*
 * @lc app=leetcode.cn id=275 lang=golang
 *
 * [275] H 指数 II
 */

// @lc code=start
func hIndex(citations []int) int {
	n := len(citations)
	left, right := 0, n-1

	for left <= right {
		mid := left+(right-left)/2

		// 如果倒数第 mid 个数字也符合要求，说明还可以变大
		if citations[n-1-mid] >= mid+1 {
			left = mid+1
		} else {
			right = mid-1
		}
	}

	return left
}
// @lc code=end

