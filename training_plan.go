package main

// https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
func trainingPlan(actions []int) []int {
	if len(actions) == 0 {
		return actions
	}

	var prev, current int

	for current < len(actions) {
		if actions[prev]%2 == 0 && actions[current]%2 != 0 {
			actions[prev], actions[current] = actions[current], actions[prev]
			prev++
		} else if actions[prev]%2 != 0 {
			prev++
		}

		current++
	}

	return actions
}
