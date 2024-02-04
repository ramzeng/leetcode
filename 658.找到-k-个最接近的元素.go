/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] 找到 K 个最接近的元素
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
	sort.SliceStable(arr, func(i, j int) bool {
		return abs(arr[i]-x) < abs(arr[j]-x)
	})
	arr = arr[:k]
	sort.Ints(arr)
	return arr
	// right := sort.SearchInts(arr, x)
	// left := right - 1
	// n := len(arr)

	// for ; k > 0; k-- {
	// 	if left < 0 {
	// 		right++
	// 	} else if right >= n || abs(arr[left]-x) <= abs(arr[right]-x) {
	// 		left--
	// 	} else {
	// 		right++
	// 	}
	// }

	// return arr[left+1 : right]
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// @lc code=end

