#
# @lc app=leetcode.cn id=658 lang=python
#
# [658] 找到 K 个最接近的元素
#

# @lc code=start
from bisect import bisect_left

class Solution(object):
    def findClosestElements(self, arr, k, x):
        """
        :type arr: List[int]
        :type k: int
        :type x: int
        :rtype: List[int]
        """
        # 方法一
        # 先自定义排序，abs(v-x) 作为排序规则，升序排序
        # 取前 k 个元素
        # 再次升序排序
        # return sorted(sorted(arr, key=lambda v: abs(v-x))[:k])

        # 方法二
        # 二分找到下标，满足 nums[i-1] < x <= nums[i]，所以需要调用 bisect_left 函数
        right = bisect_left(arr, x)
        left = right-1
        n = len(arr)
        for _ in range(0, k):
            if left < 0:
                right += 1
            elif right >= n or abs(arr[left]-x) <= abs(arr[right]-x):
                left -= 1
            else:
                right += 1
        return arr[left+1:right]
# @lc code=end

