package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
}

// https://leetcode.cn/problems/maximum-product-of-word-lengths
func maxProduct(words []string) int {
	var masks []int
	var answer int

	// 预处理，加快下面的比较过程
	for _, word := range words {
		var mask int

		// 例如：c 的 ascii 是 99, a 是97
		// 1 << (char - 'a') 可以得到 0100
		// 遍历 word 的每一个 char，得到最终的 mask
		// 两个 word 的 mask 按位与，如果结果为 0，则一定不包含相同字符
		for _, char := range word {
			mask |= 1 << (char - 'a')
		}

		masks = append(masks, mask)
	}

	for left, _ := range masks {
		for right := left + 1; right < len(masks); right++ {
			if masks[left]&masks[right] == 0 && len(words[left])*len(words[right]) > answer {
				answer = len(words[left]) * len(words[right])
			}
		}
	}

	return answer
}
