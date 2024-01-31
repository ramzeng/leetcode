/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
var (
	dx = []int{1, 0, -1, 0}
	dy = []int{0, 1, 0, -1}
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	var dfs func(grid [][]byte, row, col, m, n int)
	dfs = func(grid [][]byte, row, col, m, n int) {
		if row >= m || col >= n || row < 0 || col < 0 {
			return
		}

		if grid[row][col] != '1' {
			return
		}

		grid[row][col] = '0'

		for i := 0; i < 4; i++ {
			dfs(grid, row+dx[i], col+dy[i], m, n)
		}
	}

	var count int
	m, n := len(grid), len(grid[0])

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] == '1' {
				count++
				dfs(grid, row, col, m, n)
			}
		}
	}

	return count
}

// @lc code=end

