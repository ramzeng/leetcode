/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	viewed := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		v := target - nums[i]
		if j, ok := viewed[v]; ok {
			return []int{j, i}
		} else {
			viewed[nums[i]] = i
		}
	}

	return nil
}

// @lc code=end

