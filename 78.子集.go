/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	var answers [][]int
	var answer []int
	var dfs func(i int)

	n := len(nums)

	dfs = func(i int) {
		answers = append(answers, append([]int{}, answer...))

		if i == n {
			return
		}

		for j := i; j < n; j++ {
			answer = append(answer, nums[j])
			dfs(j+1)
			answer = answer[:len(answer)-1]
		}
	}

	dfs(0)

	return answers
}
// @lc code=end

