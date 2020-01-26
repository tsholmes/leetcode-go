"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
"""
class Solution(object):
    def treeToDoublyList(self, root):
        """
        :type root: Node
        :rtype: Node
        """
        if root is None:
            return None

        left, right = self.walk(root)
        left.left = right
        right.right = left
        return left

    def walk(self, node):
        left, right = node, node
        if node.left is not None:
            ll, lr = self.walk(node.left)
            left = ll
            lr.right = node
            node.left = lr
        if node.right is not None:
            rl, rr = self.walk(node.right)
            right = rr
            node.right = rl
            rl.left = node
        return left, right
