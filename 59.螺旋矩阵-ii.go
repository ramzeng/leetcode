/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	rowsCount, colsCount := n, n
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	index := 1

	// 边界
	left, top, right, bottom := 0, 0, colsCount-1, rowsCount-1

	// 左边小于等于右边，上边小于等于下边
	for left <= right && top <= bottom {
		// 从左边开始遍历到右边
		for col := left; col <= right; col++ {
			matrix[top][col] = index
			index++
		}

		// 从上边遍历到下边
		for row := top + 1; row <= bottom; row++ {
			matrix[row][right] = index
			index++
		}

		// 保证至少还有两行两列
		if left < right && top < bottom {
			// 从右边遍历到左边
			for col := right - 1; col > left; col-- {
				matrix[bottom][col] = index
				index++
			}

			// 从下边遍历到上边
			for row := bottom; row > top; row-- {
				matrix[row][left] = index
				index++
			}
		}

		left++
		right--
		top++
		bottom--
	}

	return matrix
}

// @lc code=end

