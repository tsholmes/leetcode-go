"""
# Definition for a Node.
class Node(object):
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""
class Solution(object):
    def levelOrder(self, root):
        """
        :type root: Node
        :rtype: List[List[int]]
        """
        res = []
        level = []
        if root is not None:
            level.append(root)
        while len(level) > 0:
            res.append([n.val for n in level])
            level = [n2 for n in level for n2 in (n.children or [])]
        return res
