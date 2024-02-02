/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, -1, 0, 1}

	rowsCount, colsCount := len(grid), len(grid[0])

	var dfs func(grid [][]byte, row, col, rowsCount, colsCount int)
	dfs = func(grid [][]byte, row, col, rowsCount, colsCount int) {
		if row >= rowsCount || col >= colsCount {
			return
		}

		if row < 0 || col < 0 || grid[row][col] != '1' {
			return
		}

		grid[row][col] = 0

		for i := 0; i < 4; i++ {
			dfs(grid, row+dx[i], col+dy[i], rowsCount, colsCount)
		}
	}

	islandsCount := 0

	for i := 0; i < rowsCount; i++ {
		for j := 0; j < colsCount; j++ {
			if grid[i][j] == '1' {
				islandsCount++
				dfs(grid, i, j, rowsCount, colsCount)
			}
		}
	}

	return islandsCount
}

// @lc code=end

