/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	k := x ^ y

	count := 0
	for ; k > 0; k -= lowbit(k) {
		count++
	}
	return count
}

func lowbit(x int) int {
	return x & (-x)
}

// @lc code=end

