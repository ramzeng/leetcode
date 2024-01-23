/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	// p
	// s-p
	// p-(s-p)=target
	// p = s+target/2

	for _, num := range nums {
		target += num
	}

	if target < 0 || target%2 != 0 {
		return 0
	}

	target /= 2

	answers := make([]int, target+1)
	answers[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			answers[j] = answers[j] + answers[j-nums[i]]
		}
	}

	return answers[target]
}

// @lc code=end

