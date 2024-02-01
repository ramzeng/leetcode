#
# @lc app=leetcode.cn id=22 lang=python
#
# [22] 括号生成
#

# @lc code=start
class Solution(object):
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        paths = []
        path = []
        def dfs(left, right):
            if left == n and right == n:
                paths.append("".join(path))
                return
            if left < n:
                path.append("(")
                dfs(left+1, right)
                path.pop()
            if right < left:
                path.append(")")
                dfs(left, right+1)
                path.pop()
        dfs(0, 0)
        return paths
# @lc code=end

