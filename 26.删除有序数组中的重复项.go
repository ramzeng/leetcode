/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	// 快慢指针也能做
	// 数组，快慢指针
	var slow, fast int

	for ; fast < len(nums); fast++ {
		if nums[fast] == nums[slow] {
			continue
		}

		slow++
		nums[slow] = nums[fast]
	}

	// slow 是下标，答案需要是长度，所以要 +1
	return slow + 1

	// 数组题，删除有序数组中的重复项，保留 0 个，模版题
	// f := func(k int) int {
	// 	u := 0
	// 	for i := 0; i < len(nums); i++ {
	// 		if u < k || nums[u-k] != nums[i] {
	// 			nums[u] = nums[i]
	// 			u++
	// 		}
	// 	}

	// 	return u
	// }

	// return f(1)
}

// @lc code=end

