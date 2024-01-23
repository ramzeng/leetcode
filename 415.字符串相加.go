/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	var sum, carry int
	var i, j int
	var answer string

	i, j = len(num1)-1, len(num2)-1

	for ; i >= 0 || j >= 0; {
		var x, y int

		if i >= 0 {
			x = int(num1[i] - '0')
			i--
		}

		if j >= 0 {
			y = int(num2[j] - '0')
			j--
		}

		sum = x+y+carry
		sum, carry = sum%10, sum/10

		answer = strconv.Itoa(sum)+answer
	}

	if carry > 0 {
		answer = strconv.Itoa(carry)+answer
	}

	return answer
}
// @lc code=end

