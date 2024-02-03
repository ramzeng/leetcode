/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])

	i, j := 0, n-1

	for i < m && j >= 0 {
		num := matrix[i][j]
		if num == target {
			return true
		} else if num < target {
			i++
		} else {
			j--
		}
	}

	return false
}

// @lc code=end

