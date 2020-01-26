"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, prev, next, child):
        self.val = val
        self.prev = prev
        self.next = next
        self.child = child
"""
class Solution(object):
    def flatten(self, head):
        if head is None:
            return None
        """
        :type head: Node
        :rtype: Node
        """
        def walk(n, res):
            res.append(n)
            if n.child is not None:
                walk(n.child, res)
            if n.next is not None:
                walk(n.next, res)
        res = []
        walk(head, res)
        for i in range(len(res)-1):
            res[i].next = res[i+1]
            res[i+1].prev = res[i]
        res[0].prev = None
        res[-1].next = None
        for n in res:
            n.child = None
        return res[0]
