/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	var paths []string
	var path []string
	var dfs func(i int)

	n := len(s)

	dfs = func(i int) {
		if len(path) > 4 || i > n {
			return
		}

		if len(path) == 4 && i == n {
			paths = append(paths, strings.Join(path, "."))
			return
		}

		for j := i; j < n; j++ {
			if isValidIP(s[i : j+1]) {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)

	return paths
}

func isValidIP(s string) bool {
	if len(s) == 0 || len(s) > 3 {
		return false
	}

	num, _ := strconv.Atoi(s)

	if num > 255 || (s[0] == '0' && len(s) > 1) {
		return false
	}

	return true
}

// @lc code=end

