# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:   
    def isSubtree(self, root: Optional[TreeNode], subroot: Optional[TreeNode]) -> bool:
        if not subroot:
            return True
        if not root:
            return False
        if self.isSameTree(root, subroot):
            return True
        return (self.isSubtree(root.left, subroot) or
                 self.isSubtree(root.right, subroot))
    def isSameTree(self, r, s):
        if not r and not s:
            return True
        if not r or not s or r.val != s.val:
            return False
        return (self.isSameTree(r.left, s.left) and
                 self.isSameTree(r.right, s.right))
