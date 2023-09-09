package main

// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof
func reverseLeftWords(s string, n int) string {
	if len(s) == 0 {
		return s
	}

	appendString := s[:n]
	s = s[n:]

	return s + appendString
}
