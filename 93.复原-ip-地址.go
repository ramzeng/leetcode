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
		if len(path) == 4 && i == n {
			paths = append(paths, strings.Join(path, "."))
			return
		}

		for j := i; j < n; j++ {
			if isValidIPSegment(s[i : j+1]) {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)

	return paths
}

func isValidIPSegment(s string) bool {
	if len(s) > 3 || len(s) == 0 {
		return false
	}

	if len(s) > 1 && s[0] == '0' {
		return false
	}

	num, _ := strconv.Atoi(s)

	if num > 255 {
		return false
	}

	return true
}

// @lc code=end

