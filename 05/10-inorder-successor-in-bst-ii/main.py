"""
# Definition for a Node.
class Node:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None
        self.parent = None
"""
class Solution(object):
    def inorderSuccessor(self, node):
        """
        :type node: Node
        :rtype: Node
        """
        if node is None:
            return None

        if node.right is None:
            parent = node.parent
            while parent is not None:
                if parent.left == node:
                    return parent
                node, parent = parent, parent.parent
            return parent

        node = node.right
        while node.left is not None:
            node = node.left
        return node
