/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	left, right := 0, x

	for left <= right {
		mid := left + (right-left)/2

		if mid*mid >= x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left*left == x {
		return left
	} else {
		return right
	}
}

// @lc code=end

