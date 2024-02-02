/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	store := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if _, ok := store[s[i]]; ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			bracket := stack[len(stack)-1]
			if store[bracket] != s[i] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}

// @lc code=end

