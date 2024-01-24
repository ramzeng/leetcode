/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	// 快速排序
	// var quickSort func(nums []int, left, right int)
	// quickSort = func(nums []int, left, right int) {
	// 	if left >= right {
	// 		return
	// 	}

	// 	mid := nums[left+(right-left)/2]
	// 	i, j := left, right

	// 	for i < j {
	// 		for nums[i] < mid {
	// 			i++
	// 		}

	// 		for nums[j] > mid {
	// 			j--
	// 		}

	// 		if i <= j {
	// 			nums[i], nums[j] = nums[j], nums[i]
	// 			i++
	// 			j--
	// 		}
	// 	}

	// 	quickSort(nums, left, j)
	// 	quickSort(nums, i, right)
	// }

	// quickSort(nums, 0, len(nums)-1)

	// return nums

	// 归并排序
	var mergeSort func(nums []int, left, right int)
	mergeSort = func(nums []int, left, right int) {
		if left >= right {
			return
		}

		mid := left + (right-left)/2
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)

		i, j := left, mid+1
		values := make([]int, right-left+1)
		var k int

		for ; i <= mid && j <= right; k++ {
			if nums[i] < nums[j] {
				values[k] = nums[i]
				i++
			} else {
				values[k] = nums[j]
				j++
			}
		}

		for ; i <= mid; k++ {
			values[k] = nums[i]
			i++
		}

		for ; j <= right; k++ {
			values[k] = nums[j]
			j++
		}

		for i, v := range values {
			nums[left+i] = v
		}
	}

	mergeSort(nums, 0, len(nums)-1)

	return nums
}

// @lc code=end

