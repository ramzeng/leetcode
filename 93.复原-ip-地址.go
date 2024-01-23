/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	var answers []string
	var answer []string
	var dfs func(i int)

	n := len(s)
	dfs = func(i int) {
		if len(answer) > 4 || i > n {
			return
		}

		if len(answer) == 4 && i == n {
			answers = append(answers, strings.Join(answer, "."))
			return
		}

		for j := i; j < n; j++ {
			if isValidIP(s[i:j+1]) {
				answer = append(answer, s[i:j+1])
				dfs(j+1)
				answer = answer[:len(answer)-1]
			}
		}
	}

	dfs(0)

	return answers
}

func isValidIP(s string) bool {
	if len(s) == 0 || len(s) > 3 {
		return false
	}

	num, err := strconv.Atoi(s)

	if err != nil {
		return false
	}

	if num > 255 || (s[0] == '0' && len(s) > 1) {
		return false
	}

	return true
}
// @lc code=end

