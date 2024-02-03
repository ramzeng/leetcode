/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	maxArea := 0
	rowsCount, colsCount := len(grid), len(grid[0])
	var dfs func(grid [][]int, row, col, rowsCount, colsCount int) int
	dfs = func(grid [][]int, row, col, rowsCount, colsCount int) int {
		if row < 0 || col < 0 || row >= rowsCount || col >= colsCount {
			return 0
		}

		if grid[row][col] != 1 {
			return 0
		}

		grid[row][col] = 0
		area := 1

		for i := 0; i < 4; i++ {
			area += dfs(grid, row+dx[i], col+dy[i], rowsCount, colsCount)
		}

		return area
	}

	for i := 0; i < rowsCount; i++ {
		for j := 0; j < colsCount; j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs(grid, i, j, rowsCount, colsCount))
			}
		}
	}

	return maxArea
}

// @lc code=end

