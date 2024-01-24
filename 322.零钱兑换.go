/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	// 完全背包，最小值问题
	// 朴素做法
	// n := len(coins)
	// cached := make([][]int, n)

	// for i, _ := range cached {
	// 	cached[i] = make([]int, amount+1)
	// 	for j, _ := range cached[i] {
	// 		cached[i][j] = -1
	// 	}
	// }

	// var dfs func(i int, amount int) int
	// dfs = func(i int, amount int) int {
	// 	if i < 0 {
	// 		if amount == 0 {
	// 			return 0
	// 		}

	// 		return math.MaxInt / 2
	// 	}

	// 	if cached[i][amount] != -1 {
	// 		return cached[i][amount]
	// 	}

	// 	var answer int

	// 	if amount < coins[i] {
	// 		answer = dfs(i-1, amount)
	// 	} else {
	// 		answer = min(dfs(i-1, amount), dfs(i, amount-coins[i])+1)
	// 	}

	// 	cached[i][amount] = answer
	// 	return answer
	// }

	// answer := dfs(n-1, amount)

	// if answer >= math.MaxInt/2 {
	// 	return -1
	// } else {
	// 	return answer
	// }

	// 递推做法
	n := len(coins)
	answers := make([]int, amount+1)

	for i, _ := range answers {
		answers[i] = math.MaxInt / 2
	}

	answers[0] = 0

	for i := 0; i < n; i++ {
		for j := coins[i]; j <= amount; j++ {
			answers[j] = min(answers[j], answers[j-coins[i]]+1)
		}
	}

	if answers[amount] == math.MaxInt/2 {
		return -1
	} else {
		return answers[amount]
	}
}

// @lc code=end

