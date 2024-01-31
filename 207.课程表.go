/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	in := make([]int, numCourses)
	g := make([][]int, numCourses)

	// 初始化空间
	for i := 0; i < numCourses; i++ {
		g[i] = make([]int, 0)
	}

	// 遍历依赖关系，初始化入度和临接表
	for _, value := range prerequisites {
		// 前置课程，当前课程
		prev, current := value[1], value[0]
		// 当前课程入度 +1
		in[current]++
		// 更新前置课程临接表
		g[prev] = append(g[prev], current)
	}

	// bfs
	var queue []int

	for i := 0; i < numCourses; i++ {
		// 所有入度为 0 的课程进队列
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}

	var courses []int

	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		courses = append(courses, i)

		// 第 i 节课被选了，所以要根据临接表更新有依赖的课程的入度
		for _, current := range g[i] {
			in[current]--

			if in[current] == 0 {
				queue = append(queue, current)
			}
		}
	}

	return numCourses == len(courses)
}

// @lc code=end

