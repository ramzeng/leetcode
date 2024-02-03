/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	var quickSort func(nums []int, left, right int)
	quickSort = func(nums []int, left, right int) {
		if left >= right {
			return
		}

		pivot := nums[left+(right-left)/2]
		i, j := left, right

		for i <= j {
			for nums[i] < pivot {
				i++
			}

			for nums[j] > pivot {
				j--
			}

			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}

		quickSort(nums, i, right)
		quickSort(nums, left, j)
	}

	quickSort(nums, 0, len(nums)-1)
}

// @lc code=end

