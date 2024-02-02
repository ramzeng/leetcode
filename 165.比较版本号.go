/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	m, n := len(version1), len(version2)
	i, j := 0, 0

	for i < m || j < n {
		x, y := 0, 0
		for i < m && version1[i] != '.' {
			x = 10*x + int(version1[i]-'0')
			i++
		}
		i++

		for j < n && version2[j] != '.' {
			y = 10*y + int(version2[j]-'0')
			j++
		}
		j++

		if x > y {
			return 1
		}

		if x < y {
			return -1
		}
	}

	return 0
}

// @lc code=end

