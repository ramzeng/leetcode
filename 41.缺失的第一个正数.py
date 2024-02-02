#
# @lc app=leetcode.cn id=41 lang=python
#
# [41] 缺失的第一个正数
#

# @lc code=start
class Solution(object):
    def firstMissingPositive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums)
        # 原地哈希
        for i in range(0, n):
            while 1 <= nums[i] <= n and nums[i] != nums[nums[i] - 1]:
                nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]

        for i in range(0, n):
            if nums[i] != i+1:
                return i+1
        return n+1
# @lc code=end

