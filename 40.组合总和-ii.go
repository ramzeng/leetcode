/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	var paths [][]int
	var path []int
	var dfs func(i int, target int)

	sort.Ints(candidates)
	n := len(candidates)
	viewed := make([]bool, n)
	dfs = func(i int, target int) {
		if target < 0 {
			return
		}

		if target == 0 {
			paths = append(paths, append([]int{}, path...))
			return
		}

		for j := i; j < n; j++ {
			if j > 0 && candidates[j-1] == candidates[j] && !viewed[j-1] {
				continue
			}

			if target < candidates[j] {
				continue
			}

			if !viewed[j] {
				path = append(path, candidates[j])
				viewed[j] = true
				dfs(j+1, target-candidates[j])
				viewed[j] = false
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0, target)

	return paths
}

// @lc code=end

