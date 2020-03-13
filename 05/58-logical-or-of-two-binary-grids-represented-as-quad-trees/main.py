"""
# Definition for a QuadTree node.
class Node(object):
    def __init__(self, val, isLeaf, topLeft, topRight, bottomLeft, bottomRight):
        self.val = val
        self.isLeaf = isLeaf
        self.topLeft = topLeft
        self.topRight = topRight
        self.bottomLeft = bottomLeft
        self.bottomRight = bottomRight
"""
class Solution(object):
    def intersect(self, quadTree1, quadTree2):
        """
        :type quadTree1: Node
        :type quadTree2: Node
        :rtype: Node
        """
        if quadTree1.isLeaf:
            if quadTree1.val == 1:
                return quadTree1
            else:
                return quadTree2
        if quadTree2.isLeaf:
            if quadTree2.val == 1:
                return quadTree2
            else:
                return quadTree1

        n = Node(
            val=0,
            isLeaf=False,
            topLeft=self.intersect(quadTree1.topLeft, quadTree2.topLeft),
            topRight=self.intersect(quadTree1.topRight, quadTree2.topRight),
            bottomLeft=self.intersect(quadTree1.bottomLeft, quadTree2.bottomLeft),
            bottomRight=self.intersect(quadTree1.bottomRight, quadTree2.bottomRight)
        )
        if n.topLeft.isLeaf and n.topRight.isLeaf and n.bottomLeft.isLeaf and n.bottomRight.isLeaf and \
            n.topLeft.val == n.topRight.val and n.bottomLeft.val == n.bottomRight.val and n.topLeft.val == n.bottomLeft.val:
            return Node(n.topLeft.val, True, None, None, None, None)
        return n
