/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	store := []string{
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

	n := len(digits)

	var paths []string
	var path []byte
	var dfs func(i int)

	dfs = func(i int) {
		if i == n {
			paths = append(paths, string(path))
			return
		}

		letters := store[int(digits[i]-'0')]

		for _, letter := range letters {
			path = append(path, byte(letter))
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return paths
}

// @lc code=end

