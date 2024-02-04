/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {
	var paths [][]string
	var path []string
	var dfs func(i int)

	n := len(s)

	dfs = func(i int) {
		if i == n {
			paths = append(paths, append([]string{}, path...))
			return
		}

		for j := i; j < n; j++ {
			if isPalindrome(s[i : j+1]) {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)

	return paths
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

// @lc code=end

