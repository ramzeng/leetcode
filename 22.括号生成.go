/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	// 回溯，组合问题
	var paths []string
	var path []byte
	var dfs func(left, right int)

	dfs = func(left, right int) {
		// 超过边界，直接退出
		if left > n || right > n {
			return
		}

		// 边界情况，收集答案
		if left == n && right == n {
			paths = append(paths, string(path))
			return
		}

		// 选左括号
		if left < n {
			path = append(path, '(')
			dfs(left+1, right)
			path = path[:len(path)-1]
		}

		// 选右括号
		if right < left {
			path = append(path, ')')
			dfs(left, right+1)
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0)

	return paths
}

// @lc code=end

