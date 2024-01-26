/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	// 回溯，组合问题，重复数组，不允许出现重复答案
	// 排序，让相同的数字连在一起
	sort.Ints(nums)

	var paths [][]int
	var path []int
	var dfs func(i int)

	n := len(nums)

	// 记录使用过的下标
	viewed := make([]bool, n)

	dfs = func(i int) {
		// 边界条件，收集答案
		if len(path) == n {
			paths = append(paths, append([]int{}, path...))
			return
		}

		// 从答案出发
		for j := 0; j < n; j++ {
			// j > 0 且当前数字等于上一个数字且上一个数字没有被使用
			if j > 0 && nums[j] == nums[j-1] && !viewed[j-1] {
				continue
			}

			if !viewed[j] {
				path = append(path, nums[j])
				viewed[j] = true
				dfs(i + 1)
				viewed[j] = false
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)

	return paths
}

// @lc code=end

