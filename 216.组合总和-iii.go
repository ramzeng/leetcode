/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	var paths [][]int
	var path []int
	var dfs func(i int, target int)

	dfs = func(i int, target int) {
		if len(path) > k {
			return
		}

		if target == 0 && len(path) == k {
			paths = append(paths, append([]int{}, path...))
			return
		}

		for j := i; j < 10; j++ {
			path = append(path, j)
			dfs(j+1, target-j)
			path = path[:len(path)-1]
		}
	}

	dfs(1, n)

	return paths
}

// @lc code=end

