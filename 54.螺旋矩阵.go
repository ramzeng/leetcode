/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	rowsCount, colsCount := len(matrix), len(matrix[0])
	path := make([]int, rowsCount*colsCount)
	index := 0
	left, right, top, bottom := 0, colsCount-1, 0, rowsCount-1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			path[index] = matrix[top][i]
			index++
		}

		for i := top + 1; i <= bottom; i++ {
			path[index] = matrix[i][right]
			index++
		}

		if left < right && top < bottom {
			for i := right - 1; i >= left; i-- {
				path[index] = matrix[bottom][i]
				index++
			}

			for i := bottom - 1; i > top; i-- {
				path[index] = matrix[i][left]
				index++
			}
		}

		left++
		right--
		top++
		bottom--
	}

	return path
}

// @lc code=end

