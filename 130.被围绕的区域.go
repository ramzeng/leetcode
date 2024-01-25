/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
var (
	dx = []int{1, 0, -1, 0}
	dy = []int{0, 1, 0, -1}
)

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	row := len(board)
	col := len(board[0])

	var dfs func(board [][]byte, i, j, row, col int)
	dfs = func(board [][]byte, i, j, row, col int) {
		if i < 0 || j < 0 || i >= row || j >= col || board[i][j] != 'O' {
			return
		}

		board[i][j] = 'A'

		for t := 0; t < 4; t++ {
			dfs(board, i+dx[t], j+dy[t], row, col)
		}
	}

	for i := 0; i < row; i++ {
		// 第 i 行，第 1 列
		dfs(board, i, 0, row, col)
		// 第 i 行，最后一列
		dfs(board, i, col-1, row, col)
	}

	for i := 0; i < col; i++ {
		dfs(board, 0, i, row, col)
		dfs(board, row-1, i, row, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

// @lc code=end

