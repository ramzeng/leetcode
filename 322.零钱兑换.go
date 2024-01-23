/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	answers := make([]int, amount+1)

	for i := 0; i < amount+1; i++ {
		answers[i] = amount + 1
	}

	n := len(coins)

	answers[0] = 0

	for i := 0; i < n; i++ {
		for j := coins[i]; j <= amount; j++ {
			answers[j] = min(answers[j], answers[j-coins[i]]+1)
		}
	}

	if answers[amount] == amount+1 {
		return -1
	}

	return answers[amount]
}

// @lc code=end

