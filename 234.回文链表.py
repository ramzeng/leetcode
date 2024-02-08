#
# @lc app=leetcode.cn id=234 lang=python3
#
# [234] 回文链表
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        mid = self.findMiddleNode(head)
        p = mid.next
        mid.next = None
        p = self.reverseNodeList(p)

        while head and p:
            if head.val != p.val:
                return False
            head = head.next
            p = p.next
        return True

    def findMiddleNode(self, head: Optional[ListNode]) -> ListNode:
        slow, fast = head, head.next
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
        return slow

    def reverseNodeList(self, head: Optional[ListNode]) -> ListNode:
        prev, current = None, head
        while current:
            next = current.next
            current.next = prev
            prev = current
            current = next
        return prev

# @lc code=end

