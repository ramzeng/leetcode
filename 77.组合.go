/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	var paths [][]int
	var path []int
	var dfs func(i int)

	dfs = func(i int) {
		d := k - len(path)

		if d == 0 {
			paths = append(paths, append([]int{}, path...))
			return
		}

		for j := i; j <= n; j++ {
			path = append(path, j)
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(1)

	return paths
}

// @lc code=end

