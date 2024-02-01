#
# @lc app=leetcode.cn id=165 lang=python
#
# [165] 比较版本号
#

# @lc code=start
class Solution(object):
    def compareVersion(self, version1, version2):
        """
        :type version1: str
        :type version2: str
        :rtype: int
        """
        m, n = len(version1), len(version2)
        i, j = 0, 0

        while i < m or j < n:
            x, y = 0, 0
            while i < m and version1[i] != ".":
                x = x*10+int(version1[i])
                i += 1
            i += 1
            while j < n and version2[j] != ".":
                y = y*10+int(version2[j])
                j += 1
            j += 1
            if x > y:
                return 1
            elif x < y:
                return -1
        return 0
# @lc code=end

