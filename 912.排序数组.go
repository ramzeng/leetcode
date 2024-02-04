/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	// var quickSort func(nums []int, left, right int)
	// quickSort = func(nums []int, left, right int) {
	// 	if left >= right {
	// 		return
	// 	}

	// 	pivot := nums[left+(right-left)/2]
	// 	i, j := left, right

	// 	for i <= j {
	// 		for nums[i] < pivot {
	// 			i++
	// 		}
	// 		for nums[j] > pivot {
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

	var mergeSort func(nums []int, left, right int)
	mergeSort = func(nums []int, left, right int) {
		if left >= right {
			return
		}

		mid := left + (right-left)/2
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)

		i, j := left, mid+1
		k := 0
		sortedNums := make([]int, right-left+1)

		for ; i <= mid && j <= right; k++ {
			if nums[i] < nums[j] {
				sortedNums[k] = nums[i]
				i++
			} else {
				sortedNums[k] = nums[j]
				j++
			}
		}

		for ; i <= mid; k++ {
			sortedNums[k] = nums[i]
			i++
		}

		for ; j <= right; k++ {
			sortedNums[k] = nums[j]
			j++
		}

		for i := 0; i < k; i++ {
			nums[left+i] = sortedNums[i]
		}
	}

	mergeSort(nums, 0, len(nums)-1)
	return nums
}

// @lc code=end

