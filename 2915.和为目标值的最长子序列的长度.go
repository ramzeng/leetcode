/*
 * @lc app=leetcode.cn id=2915 lang=golang
 *
 * [2915] 和为目标值的最长子序列的长度
 */

// @lc code=start
func lengthOfLongestSubsequence(nums []int, target int) int {
	// 0 / 1 背包，最大值
	n := len(nums)

	// 缓存结果
	cached := make([][]int, n)

	for i, _ := range cached {
		cached[i] = make([]int, target+1)

		for j, _ := range cached[i] {
			// 从前 i 个选，和为 j 的最大被选数字数量
			cached[i][j] = -1
		}
	}

	var dfs func(i int, target int) int
	dfs = func(i int, target int) int {
		if i < 0 {
			if target == 0 {
				return 0
			}

			return math.MinInt
		}

		if cached[i][target] != -1 {
			return cached[i][target]
		}

		var answer int

		if target < nums[i] {
			answer = dfs(i-1, target)
		} else {
			answer = max(dfs(i-1, target), dfs(i-1, target-nums[i])+1)
		}

		cached[i][target] = answer
		return answer
	}

	answer := dfs(n-1, target)

	if answer < 0 {
		return -1
	} else {
		return answer
	}
}

// @lc code=end

