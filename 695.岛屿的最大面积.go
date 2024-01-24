/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	// DFS 朴素解法
	var dfs func(grid [][]int, i, j, row, col int) int
	dfs = func(grid [][]int, i, j, row, col int) int {
		if i < 0 || j < 0 || i >= row || j >= col || grid[i][j] == 0 {
			return 0
		}

		grid[i][j] = 0

		return dfs(grid, i-1, j, row, col) + dfs(grid, i+1, j, row, col) + dfs(grid, i, j-1, row, col) + dfs(grid, i, j+1, row, col) + 1
	}

	answer := 0
	row := len(grid)
	col := len(grid[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				answer = max(answer, dfs(grid, i, j, row, col))
			}
		}
	}

	return answer
}

// @lc code=end

