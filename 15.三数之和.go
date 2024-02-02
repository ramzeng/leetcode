/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	n := len(nums)
	sum := 0
	values := [][]int{}

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

		for left < right {
			sum = nums[i]+nums[left]+nums[right]
			if sum == 0 {
				values = append(values, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for right > left && nums[right] == nums[right-1] {
					right--
				}
			}

			sum = nums[i]+nums[left]+nums[right]

			if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return values
}
// @lc code=end

