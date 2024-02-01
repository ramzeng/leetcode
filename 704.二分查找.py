#
# @lc app=leetcode.cn id=704 lang=python
#
# [704] 二分查找
#

# @lc code=start
class Solution(object):
    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        # 标准库
        # left = bisect_left(nums, target)
        # if left == len(nums) or nums[left] != target:
        #     return -1
        # else:
        #     return left

        # 朴素实现
        n = len(nums)
        left, right = 0, n-1

        while left <= right:
            mid = left+(right-left)//2
            if nums[mid] >= target:
                right = mid-1
            else:
                left = mid+1
        if left == n or nums[left] != target:
            return -1
        else:
            return left
# @lc code=end

