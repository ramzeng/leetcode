/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	// 排序，让相同的元素贴在一起
	sort.Ints(candidates)

	// 回溯，组合问题
	var paths [][]int
	var path []int
	var dfs func(i int, target int)

	n := len(candidates)

	dfs = func(i int, target int) {
		// 边界情况，收集答案
		if target == 0 {
			paths = append(paths, append([]int{}, path...))
			return
		}

		// 从答案出发
		for j := i; j < n; j++ {
			// 跳过重复的元素
			// j > i，说明开始 dfs 了
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}

			// 判断是否可以选这个数
			if target >= candidates[j] {
				path = append(path, candidates[j])
				// 可以重复选的情况
				// dfs(j, target-candidates[j])
				// 不可以重复选要用下面这种
				dfs(j+1, target-candidates[j])
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0, target)

	return paths
}

// @lc code=end

