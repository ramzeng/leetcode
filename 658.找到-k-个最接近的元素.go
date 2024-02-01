/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] 找到 K 个最接近的元素
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
	sort.SliceStable(arr, func(current, next int) bool {
		return abs(arr[current]-x) < abs(arr[next]-x)
	})

	arr = arr[:k]

	sort.Ints(arr)

	return arr
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// @lc code=end

