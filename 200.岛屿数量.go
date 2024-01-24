/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
var (
	dx = [4]int{1, 0, -1, 0}
	dy = [4]int{0, 1, 0, -1}
)

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

		for x := 0; x < 4; x++ {
			dfs(grid, i+dx[x], j+dy[x], row, col)
		}
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

