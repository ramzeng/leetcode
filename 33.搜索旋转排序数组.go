/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	// 数组，二分
	// 先找到最小值的下标
	n := len(nums)
	i := findMin(nums)

	// 判断 target 与最后一个数的大小，决定二分的范围
	var left, right int
	if target <= nums[n-1] {
		left, right = i, n-1
	} else {
		left, right = 0, i-1
	}

	// 常规二分
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left == n || nums[left] != target {
		return -1
	} else {
		return left
	}
}

func findMin(nums []int) int {
	n := len(nums)
	// 这里要注意，最后一个值要作为参考值，不能包含在二分范围内
	left, right := 0, n-2

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= nums[n-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

// @lc code=end

