/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var paths [][]int
	var path []int
	var dfs func(i int, target int)

	n := len(candidates)
	dfs = func(i int, target int) {
		if target < 0 {
			return
		}

		if target == 0 {
			paths = append(paths, append([]int{}, path...))
			return
		}

		for j := i; j < n; j++ {
			if target >= candidates[j] {
				path = append(path, candidates[j])
				dfs(j, target-candidates[j])
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0, target)

	return paths
}

// @lc code=end

