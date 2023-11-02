package main

// 双指针
// 接雨水 https://leetcode.cn/problems/trapping-rain-water
func trap(height []int) int {
  left := 0
  right := len(height)-1
  prefixMax := 0
  suffixMax := 0
  answer := 0

  for left <= right {
    prefixMax = max(prefixMax, height[left])
    suffixMax = max(suffixMax, height[right])

    if prefixMax < suffixMax {
      answer += prefixMax - height[left]
      left++
    } else {
      answer += suffixMax - height[right]
      right--
    }
  }

  return answer
}
