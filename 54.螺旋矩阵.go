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

	index := 0
	rowsCount, colsCount := len(matrix), len(matrix[0])
	values := make([]int, rowsCount*colsCount)
	left, top, right, bottom := 0, 0, colsCount-1, rowsCount-1

	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			values[index] = matrix[top][col]
			index++
		}

		for row := top+1; row <= bottom; row++ {
			values[index] = matrix[row][right]
			index++
		}

		if left < right && top < bottom {
			for col := right-1; col > left; col-- {
				values[index] = matrix[bottom][col]
				index++
			}

			for row := bottom; row > top; row-- {
				values[index] = matrix[row][left]
				index++
			}
		}

		left++
		top++
		right--
		bottom--
	}

	return values
}

// @lc code=end

