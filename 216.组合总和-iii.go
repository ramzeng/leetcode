/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	var paths [][]int
	var path []int
	var dfs func(i int)

	dfs = func(i int) {
		d := k-len(path)

		if i < d || d < 0 {
			return
		}

		if d == 0 && sum(path) == n {
			paths = append(paths, append([]int{}, path...))
			return
		}

		for j := i; j >= 1; j-- {
			path = append(path, j)
			dfs(j-1)
			path = path[:len(path)-1]
		}
	}

	dfs(9)

	return paths
}

func sum(path []int) int {
	sum := 0
	for _, v := range path {
		sum += v
	}

	return sum
}

// @lc code=end

