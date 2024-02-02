/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	i, j := m-1, n-1
	sum, carry := 0, 0
	answer := ""

	for i >= 0 || j >= 0 {
		x, y := 0, 0

		if i >= 0 {
			x = int(num1[i] - '0')
			i--
		}

		if j >= 0 {
			y = int(num2[j] - '0')
			j--
		}

		sum = x + y + carry
		sum, carry = sum%10, sum/10
		answer = strconv.Itoa(sum) + answer
	}

	if carry > 0 {
		answer = strconv.Itoa(carry) + answer
	}

	return answer
}

// @lc code=end

