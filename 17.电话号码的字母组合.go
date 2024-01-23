/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var answers []string
	var answer []byte
	var dfs func(i int)

	n := len(digits)
	mapping := []string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}

	dfs = func(i int) {
		if i == n {
			answers = append(answers, string(answer))
			return
		}

		for _, r := range mapping[int(digits[i]-'0')] {
			answer = append(answer, byte(r))
			dfs(i+1)
			answer = answer[:len(answer)-1]
		}
	}

	dfs(0)

	return answers
}
// @lc code=end

