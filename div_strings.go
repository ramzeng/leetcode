package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(divStrings("21", "7"))
}

func divStrings(num1 string, num2 string) string {
	var answer string
	var div int

	x, _ := strconv.Atoi(num2)

	for left := 0; left < len(num1); left++ {
		num := int(num1[left] - '0')

		sum := div*10 + num
		sum, div = sum/x, sum%x

		answer += strconv.Itoa(sum)
	}

	answer = strings.TrimLeft(answer, "0")

	return answer
}
