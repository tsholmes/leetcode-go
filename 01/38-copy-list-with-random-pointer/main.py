# no go submissions :(
# Definition for a Node.
class Node(object):
    def __init__(self, val, next, random):
        self.val = val
        self.next = next
        self.random = random

class Solution(object):
    def copyRandomList(self, head):
        """
        :type head: Node
        :rtype: Node
        """
        if head is None:
            return None
        fwd = {}
        h = head
        while h is not None:
            h2 = Node(h.val, None, None)
            fwd[id(h)] = h2
            h = h.next

        h = head
        while h is not None:
            h2 = fwd[id(h)]
            h2.next = None if h.next is None else fwd[id(h.next)]
            h2.random = None if h.random is None else fwd[id(h.random)]
            h = h.next

        return fwd[id(head)]

if __name__ == '__main__':
    n1, n2 = Node(1, None, None), Node(1, None, None)
    n1.next, n1.random = n2, n2
    n2.random = n2
    sol = Solution()
    sol.copyRandomList(n1)
