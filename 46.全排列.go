/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	var answers [][]int
	var answer []int
	var dfs func(i int)

	n := len(nums)
	viewed := make([]bool, n)

	dfs = func(i int) {
		if len(answer) == n {
			answers = append(answers, append([]int{}, answer...))
			return
		}

		for j := 0; j < n; j++ {
			if !viewed[j] {
				answer = append(answer, nums[j])
				viewed[j] = true
				dfs(i + 1)
				viewed[j] = false
				answer = answer[:len(answer)-1]
			}
		}
	}

	dfs(0)

	return answers
}

// @lc code=end

