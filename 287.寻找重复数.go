/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	fast, slow := 0, 0
	for {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if fast == slow {
			break
		}
	}
	finder := 0
	for {
		finder = nums[finder]
		slow = nums[slow]
		if slow == finder {
			break
		}
	}
	return slow
}

// @lc code=end

