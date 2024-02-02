/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	n := len(intervals[0])
	current := intervals[0]
	answers := [][]int{}

	for _, interval := range intervals {
		left, right := interval[0], interval[n-1]

		if current[n-1] >= left {
			if current[n-1] < right {
				current[n-1] = right
			}
		} else {
			answers = append(answers, current)
			current = interval
		}
	}

	answers = append(answers, current)
	return answers
}

// @lc code=end

