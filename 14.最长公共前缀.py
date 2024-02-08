#
# @lc app=leetcode.cn id=14 lang=python3
#
# [14] 最长公共前缀
#

# @lc code=start
class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return ""

        maxStr = max(strs)
        minStr = min(strs)

        for i in range (0, len(minStr)):
            if maxStr[i] != minStr[i]:
                return minStr[:i]
        return minStr
# @lc code=end

