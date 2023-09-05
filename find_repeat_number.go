package main

import "fmt"

func main() {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}

// 输入：[2, 3, 1, 0, 2, 5, 3]
// 输出：2 或 3
// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
func findRepeatNumber(nums []int) int {
	numsCount := make(map[int]int)

	for _, num := range nums {
		if _, ok := numsCount[num]; ok {
			return num
		} else {
			numsCount[num] = 1
		}
	}

	return 0
}
