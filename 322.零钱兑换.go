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
	// dp[i][j] i 表示前 i 枚硬币，j 表示总金额，值表示从前 i 枚选的最少的硬币个数
	n := len(coins)
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 0
	}

	for i := 0; i < amount+1; i++ {
		dp[0][i] = amount + 1
	}

	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j <= amount; j++ {
			dp[i+1][j] = dp[i][j]

			if j >= coins[i] {
				// 完全背包，dp[i+1][j-coins[i]]+1
				// 01 背包, dp[i][j-coins[i]]+1
				dp[i+1][j] = min(dp[i][j], dp[i+1][j-coins[i]]+1)
			}
		}
	}

	if dp[n][amount] < amount+1 {
		return dp[n][amount]
	}

	return -1
}

// @lc code=end

