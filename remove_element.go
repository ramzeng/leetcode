package main

// 移除元素
// 双指针
// https://leetcode.cn/problems/remove-element
func removeElement(nums []int, val int) int {
	left := 0

	for _, num := range nums {
		if num != val {
			nums[left] = num
			left++
		}
	}

	return left
}
