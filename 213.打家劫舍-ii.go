/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	return max(f(nums[0:n-1]), f(nums[1:n]))
}

func f(nums []int) int {
	// 朴素做法
	// var dfs func(i int) int

	// n := len(nums)
	// cached := make([]int, n)

	// for i, _ := range cached {
	// 	cached[i] = -1
	// }

	// dfs = func(i int) int {
	// 	if i < 0 {
	// 		return 0
	// 	}

	// 	if cached[i] != -1 {
	// 		return cached[i]
	// 	}

	// 	// 3 = max(dfs(2), dfs(1)+nums[3])
	// 	// 2 = max(dfs(1), dfs(0)+nums[2])
	// 	// 1 = max(dfs(0), dfs(-1)+nums[1])
	// 	// 0 = max(dfs(-1), dfs(-2)+nums[0])
	// 	answer := max(dfs(i-1), dfs(i-2)+nums[i])
	// 	cached[i] = answer
	// 	return answer
	// }

	// return dfs(len(nums)-1)

	// 递推做法
	n := len(nums)
	answers := make([]int, n+2)

	for i := 0; i < n; i++ {
		answers[i+2] = max(answers[i+1], answers[i]+nums[i])
	}

	return answers[n+1]
}

// @lc code=end

