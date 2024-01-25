/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	// 字符串，栈
	// 定义 mapping，右括号为 K，左括号为 V
	mapping := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 借助栈结构快速判断
	var stack []byte

	for i := 0; i < len(s); i++ {
		// 如果是左括号，直接进栈即可
		if _, ok := mapping[s[i]]; !ok {
			stack = append(stack, s[i])
		} else {
			// 没有左括号匹配
			if len(stack) == 0 {
				return false
			}

			// 拿到栈顶左括号
			bracket := stack[len(stack)-1]

			// 判断是否匹配
			if mapping[s[i]] != bracket {
				return false
			}

			// 弹出
			stack = stack[:len(stack)-1]
		}
	}

	// 判断所有左括号是否都搞定了
	return len(stack) == 0
}

// @lc code=end

