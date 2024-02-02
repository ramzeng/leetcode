/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	var paths [][]int
	var path []int
	var dfs func(i int)

	n := len(nums)

	dfs = func(i int) {
		paths = append(paths, append([]int{}, path...))
		if i == n {
			return
		}

		for j := i; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return paths
}

// @lc code=end

