/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	// 0 / 1 背包，方案数问题

	// p
	// s-p
	// p-(s-p)=target
	// p = s+target/2

	// 问题转化
	// for _, num := range nums {
	// 	target += num
	// }

	// if target < 0 || target%2 != 0 {
	// 	return 0
	// }

	// target /= 2

	// n := len(nums)

	// // 朴素做法
	// cached := make([][]int, n)

	// for i, _ := range cached {
	// 	cached[i] = make([]int, target+1)

	// 	for j, _ := range cached[i] {
	// 		cached[i][j] = -1
	// 	}
	// }

	// var dfs func(i int, target int) int
	// dfs = func(i int, target int) int {
	// 	// 如果 i 超出边界，收集答案
	// 	if i < 0 {
	// 		if target == 0 {
	// 			return 1
	// 		} else {
	// 			return 0
	// 		}
	// 	}

	// 	if cached[i][target] != -1 {
	// 		return cached[i][target]
	// 	}

	// 	var answer int

	// 	// 如果 target 小于 nums[i]，只能不选
	// 	if target < nums[i] {
	// 		answer = dfs(i-1, target)
	// 	} else {
	// 		// 不选和选
	// 		answer = dfs(i-1, target)+dfs(i-1, target-nums[i])
	// 	}

	// 	cached[i][target] = answer
	// 	return answer
	// }

	// return dfs(n-1, target)

	// 递推做法
	// p
	// s-p
	// p-(s-p)=target
	// p = s+target/2

	// 问题转化
	for _, num := range nums {
		target += num
	}

	if target < 0 || target%2 != 0 {
		return 0
	}

	target /= 2

	n := len(nums)
	answers := make([]int, target+1)
	answers[0] = 1

	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			answers[j] = answers[j] + answers[j-nums[i]]
		}
	}

	return answers[target]
}

// @lc code=end

