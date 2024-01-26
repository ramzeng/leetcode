/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	// 恰好每行每列有一个
	// r 表示行号，c 表示列号
	// r+c，假设要在 (2, 0) 放置，则不能在 r+c=2 的其他地方放置，判断右上方的皇后
	// r-c，假设要在 (2, 0) 放置，则不能在 r-c=-1 的其他地方放置，判断左上方的皇后
	var boards [][]string
	var board []string
	// r 当前枚举的行号，s 表示剩余可枚举的列号
	var dfs func(row int)
	var checker func(row, col int) bool

	cols := make([]int, n)
	viewed := make([]bool, n)

	checker = func(row, col int) bool {
		for R := 0; R < row; R++ {
			C := cols[R]
			if row+col == R+C || row-col == R-C {
				return false
			}
		}

		return true
	}

	dfs = func(row int) {
		if row == n {
			board = make([]string, n)
			for i, c := range cols {
				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-1-c)
			}
			boards = append(boards, append([]string{}, board...))
			return
		}

		for col := 0; col < n; col++ {
			if !viewed[col] && checker(row, col) {
				cols[row] = col
				viewed[col] = true
				dfs(row + 1)
				viewed[col] = false
				cols[row] = 0
			}
		}
	}

	dfs(0)

	return boards
}

// @lc code=end

