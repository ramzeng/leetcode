package main

// https://leetcode.cn/problems/repeated-dna-sequences
func findRepeatedDnaSequences(s string) []string {
	var answer []string
	substrings := make(map[string]int)

	for left := 0; left < len(s)-9; left++ {
		substring := s[left : left+10]
		substrings[substring]++

		if substrings[substring] == 2 {
			answer = append(answer, substring)
		}
	}

	return answer
}
