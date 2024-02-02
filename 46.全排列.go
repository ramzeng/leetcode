/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	var paths [][]int
	var path []int
	var dfs func(i int)

	n := len(nums)
	visited := make([]bool, n)

	dfs = func(i int) {
		if i == n {
			paths = append(paths, append([]int{}, path...))
			return
		}

		for j := 0; j < n; j++ {
			if !visited[j] {
				path = append(path, nums[j])
				visited[j] = true
				dfs(i + 1)
				visited[j] = false
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)

	return paths
}

// @lc code=end

