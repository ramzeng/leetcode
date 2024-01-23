/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {
	var answers [][]string
	var answer []string
	var dfs func(i int)

	n := len(s)
	dfs = func(i int) {
		if i == n {
			answers = append(answers, append([]string{}, answer...))
			return
		}

		for j := i; j < n; j++ {
			if isPalindrome(s[i:j+1]) {
				answer = append(answer, string(s[i:j+1]))
				dfs(j+1)
				answer = answer[:len(answer)-1]
			}
		}
	}

	dfs(0)

	return answers
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

