/*
 * @lc app=leetcode.cn id=2962 lang=golang
 *
 * [2962] 统计最大元素出现至少 K 次的子数组
 */

// @lc code=start
func countSubarrays(nums []int, k int) int64 {
	maxNum := slices.Max(nums)

	var slow, fast int
	var answer, count int

	n := len(nums)

	for ; fast < n; fast++ {
		if nums[fast] == maxNum {
			count++
		}

		for count >= k {
			if nums[slow] == maxNum {
				count--
			}

			slow++
		}

		answer += slow
	}

	return int64(answer)
}

// @lc code=end

