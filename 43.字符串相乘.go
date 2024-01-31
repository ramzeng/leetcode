/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var sum, carry int
	var answer string
	y, _ := strconv.Atoi(num1)

	// 从num2的最后一位开始遍历，当i>=0或者carry>0时继续循环
	for i := len(num2) - 1; i >= 0; i-- {
		x := int(num2[i] - '0')
		sum = x*y + carry
		sum, carry = sum%10, sum/10
		answer = strconv.Itoa(sum) + answer
	}

	// 处理最后的carry
	if carry > 0 {
		answer = strconv.Itoa(carry) + answer
	}

	return answer
}

// @lc code=end

