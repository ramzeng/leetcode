package main

import (
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/largest-number/submissions/
func largestNumber(nums []int) string {
	// 转化为字符串切片
	numStrings := make([]string, len(nums))

	for k, v := range nums {
		numStrings[k] = strconv.Itoa(v)
	}

	// 降序排序
	sort.Slice(numStrings, func(i, j int) bool {
		x, y := numStrings[i], numStrings[j]

		if x+y > y+x {
			return true
		} else if x+y == y+x {
			return y < x
		} else {
			return false
		}
	})

	// 转化为字符串结果
	result := strings.Join(numStrings, "")

	// 首字节为 0 的情况
	if result[0] == '0' {
		return "0"
	}

	return result
}
