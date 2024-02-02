/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	store := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if key, ok := store[target-nums[i]]; ok {
			return []int{key, i}
		}
		store[nums[i]] = i
	}
	return nil
}

// @lc code=end

