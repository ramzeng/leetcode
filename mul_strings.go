package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(mulStrings("123", "12"))
}

func mulStrings(num1 string, num2 string) string {
	var answer string
	var carry int
	y, _ := strconv.Atoi(num1)

	for left := len(num2) - 1; left >= 0 || carry > 0; left-- {
		if left >= 0 {
			x := int(num2[left] - '0')
			carry += x * y
		}

		answer = strconv.Itoa(carry%10) + answer
		carry /= 10
	}

	return answer
}
