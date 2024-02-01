#
# @lc app=leetcode.cn id=51 lang=python
#
# [51] N 皇后
#

# @lc code=start
class Solution(object):
    def solveNQueens(self, n):
        """
        :type n: int
        :rtype: List[List[str]]
        """
        row_cols = [-1] * n
        cols = [False] * n
        paths = []
        def check(row, col):
            for R in range(0, row):
                C = row_cols[R]
                if row+col == R+C or row-col == R-C:
                    return False
            return True

        def dfs(row):
            if row > n:
                return
            if row == n:
                path = []
                for i in range(0, n):
                    col = row_cols[i]
                    path.append("."*col+"Q"+"."*(n-1-col))
                paths.append(path)
            for col in range(0, n):
                if cols[col] == False and check(row, col):
                    cols[col] = True
                    row_cols[row] = col
                    dfs(row+1)
                    row_cols[row] = -1
                    cols[col] = False
        dfs(0)
        return paths



# @lc code=end

