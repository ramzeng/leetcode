package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addStrings("11", "123"))
}

// 双指针
// https://leetcode.cn/problems/add-strings/submissions
func addStrings(num1 string, num2 string) string {
	result := ""
	overflow := 0

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var x, y int

		if i >= 0 {
			x = int(num1[i] - '0')
		}

		if j >= 0 {
			y = int(num2[j] - '0')
		}

		sum := x + y + overflow
		num := sum % 10
		overflow = sum / 10
		result = strconv.Itoa(num) + result
	}

	if overflow > 0 {
		result = strconv.Itoa(overflow) + result
	}

	return result
}
