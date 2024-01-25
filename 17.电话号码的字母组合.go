/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	// 回溯，组合问题
	if len(digits) == 0 {
		return nil
	}

	strings := []string{
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

	// 保存路径
	var path []byte
	var paths []string

	// 输入的号码长度，也是 path 的长度
	n := len(digits)

	// 深度遍历
	var dfs func(i int)
	dfs = func(i int) {
		// 边界情况，收集答案
		if i == n {
			paths = append(paths, string(path))
			return
		}

		// 从结果出发
		// i 代表在 digits 中第 i+1 个值
		// j 代表 digits[i] 映射的字符串中的第 j+1 个字符
		str := strings[int(digits[i]-'0')]
		for j := 0; j < len(str); j++ {
			path = append(path, str[j])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return paths
}

// @lc code=end

