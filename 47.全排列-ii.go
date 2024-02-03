/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var paths [][]int
	var path []int
	var dfs func(i int)

	n := len(nums)
	viewed := make([]bool, n)

	dfs = func(i int) {
		if i == n {
			paths = append(paths, append([]int{}, path...))
			return
		}

		for j := 0; j < n; j++ {
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

