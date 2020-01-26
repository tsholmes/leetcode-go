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
    def construct(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: Node
        """
        N = len(grid)
        return self.construct_range(grid, 0, N, 0, N)

    def construct_range(self, grid, r1, r2, c1, c2):
        if r2-r1 == 1:
            return Node(grid[r1][c1], True, None, None, None, None)

        p = (r2-r1)//2
        tl = self.construct_range(grid, r1, r1+p, c1, c1+p)
        tr = self.construct_range(grid, r1, r1+p, c1+p, c2)
        bl = self.construct_range(grid, r1+p, r2, c1, c1+p)
        br = self.construct_range(grid, r1+p, r2, c1+p, c2)

        ns = [tl, tr, bl, br]

        if all([n.isLeaf for n in ns]) and (all([n.val for n in ns]) == any([n.val for n in ns])):
            return Node(tl.val, True, None, None, None, None)

        return Node(None, False, tl, tr, bl, br)
