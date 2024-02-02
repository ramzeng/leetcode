/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	k = n - k
	answer := 0

	var quickSelect func(nums []int, left, right int)
	quickSelect = func(nums []int, left, right int) {
		if left >= right {
			answer = nums[k]
			return
		}

		mid := nums[left+(right-left)/2]
		i, j := left, right

		for i <= j {
			for nums[i] < mid {
				i++
			}

			for nums[j] > mid {
				j--
			}

			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}

		if k > j {
			quickSelect(nums, i, right)
		} else {
			quickSelect(nums, left, j)
		}
	}

	quickSelect(nums, 0, n-1)
	return answer
}

// @lc code=end

