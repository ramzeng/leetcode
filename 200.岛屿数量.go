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

	// DFS 朴素解法
	var dfs func(grid [][]byte, i, j, row, col int)
	dfs = func(grid [][]byte, i, j, row, col int) {
		if i < 0 || j < 0 || i >= row || j >= col || grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'

		dfs(grid, i-1, j, row, col)
		dfs(grid, i+1, j, row, col)
		dfs(grid, i, j-1, row, col)
		dfs(grid, i, j+1, row, col)
	}

	answer := 0
	row := len(grid)
	col := len(grid[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				answer++
				dfs(grid, i, j, row, col)
			}
		}
	}

	return answer
}

// @lc code=end

