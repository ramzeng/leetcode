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
	viewed := make([]bool, n)

	dfs = func(i int) {
		// 收集答案
		if len(path) == n {
			paths = append(paths, append([]int{}, path...))
			return
		}

		// 从答案出发
		for j := 0; j < n; j++ {
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

