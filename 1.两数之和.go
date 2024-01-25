/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// 数组题，求 target
	// viewed 数组记录遍历过了哪些
	// key 存数值，value 存下标
	viewed := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		value := target - nums[i]

		if j, ok := viewed[value]; ok {
			return []int{j, i}
		} else {
			viewed[nums[i]] = i
		}
	}

	return nil
}

// @lc code=end

