# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        def dfs(left, right):
            if not left and not right:
                return True
            if not left or not right or left.val != right.val:
                return False
            return (left.val == right.val and
                    dfs(left.right, right.right) and dfs(left.left, right.left))
        return dfs(p, q)