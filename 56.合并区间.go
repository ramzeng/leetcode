/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	var ranges [][]int
	var current []int

	m, n := len(intervals), len(intervals[0])

	if m <= 1 {
		return intervals
	}

	// 排序十分重要
	sort.Slice(intervals, func(current, next int) bool {
		return intervals[current][0] < intervals[next][0]
	})

	current = intervals[0]

	for i := 1; i < m; i++ {
		left, right := intervals[i][0], intervals[i][n-1]

		if current[n-1] >= left {
			if current[n-1] < right {
				current[n-1] = right
			}
		} else {
			ranges = append(ranges, current)
			current = intervals[i]
		}
	}

	ranges = append(ranges, current)

	return ranges
}

// @lc code=end

