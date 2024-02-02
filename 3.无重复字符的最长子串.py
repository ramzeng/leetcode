#
# @lc app=leetcode.cn id=3 lang=python
#
# [3] 无重复字符的最长子串
#

# @lc code=start
class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        visited = set()
        slow, max_length = 0, 0

        for fast, c in enumerate(s):
            while c in visited:
                visited.remove(s[slow])
                slow += 1
            visited.add(c)
            max_length = max(max_length, fast-slow+1)
        return max_length

# @lc code=end

