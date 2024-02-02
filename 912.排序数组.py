#
# @lc app=leetcode.cn id=912 lang=python
#
# [912] 排序数组
#

# @lc code=start
import random

class Solution(object):
    def sortArray(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        def quickSort(nums, left, right):
            if left >= right:
                return
            # 选中点，大部分情况可以，但也可能退化成 O(n2)
            # mid = nums[left+(right-left)//2]
            # 随机选取一个参考点，容错率更高
            mid = nums[random.randint(left, right)]
            i, j = left, right
            while i <= j:
                while nums[i] < mid:
                    i += 1
                while nums[j] > mid:
                    j -= 1
                if i <= j:
                    nums[i], nums[j] = nums[j], nums[i]
                    i += 1
                    j -= 1
            quickSort(nums, left, j)
            quickSort(nums, i, right)
        quickSort(nums, 0, len(nums)-1)
        return nums
# @lc code=end

