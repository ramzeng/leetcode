/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	indegrees := make([]int, numCourses)
	dependencies := make([][]int, numCourses)

	for i := 0; i < numCourses; i++ {
		dependencies[i] = make([]int, 0)
	}

	for _, prerequisit := range prerequisites {
		current, prev := prerequisit[0], prerequisit[1]
		indegrees[current]++
		dependencies[prev] = append(dependencies[prev], current)
	}

	var queue []int

	for course, indegree := range indegrees {
		if indegree == 0 {
			queue = append(queue, course)
		}
	}

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		numCourses--

		for _, course := range dependencies[course] {
			indegrees[course]--

			if indegrees[course] == 0 {
				queue = append(queue, course)
			}
		}
	}

	return numCourses == 0
}

// @lc code=end

