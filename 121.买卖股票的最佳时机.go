/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	var preMin, maxProfit int
	preMin = math.MaxInt

	for _, price := range prices {
		if price < preMin {
			preMin = price
		}

		maxProfit = max(maxProfit, price-preMin)
	}

	return maxProfit
}

// @lc code=end

