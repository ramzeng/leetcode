package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceSpace("We are happy."))
}

// 输入：s = "We are happy."
// 输出："We%20are%20happy."
// https://leetcode.cn/problems/ti-huan-kong-ge-lcof
func replaceSpace(s string) string {
	if len(s) < 1 {
		return s
	}

	result := strings.Builder{}

	for _, v := range s {
		if v == 32 {
			result.Write([]byte("%20"))
			continue
		}

		result.WriteByte(byte(v))
	}

	return result.String()
}
