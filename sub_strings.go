package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(subStrings("760", "700"))
}

func subStrings(num1 string, num2 string) string {
	var answer string
	var borrow int

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var x, y int

		if i >= 0 {
			x = int(num1[i] - '0')
		}

		if j >= 0 {
			y = int(num2[j] - '0')
		}

		borrow = x - y - borrow

		answer = strconv.Itoa((borrow+10)%10) + answer

		if borrow < 0 {
			borrow = 1
		} else {
			borrow = 0
		}
	}

	answer = strings.TrimLeft(answer, "0")

	return answer
}
