#
# @lc app=leetcode.cn id=11 lang=python3
#
# [11] 盛最多水的容器
#

# @lc code=start
class Solution:
    def maxArea(self, height: List[int]) -> int:
        n = len(height)
        left, right = 0, n-1
        maxArea, area = 0, 0
        while left < right:
            area = min(height[left], height[right]) * (right-left)
            if area > maxArea:
                maxArea = area
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        return maxArea
# @lc code=end

