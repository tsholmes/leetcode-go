# no go submissions :(
# Definition for a Node.
class Node(object):
    def __init__(self, val=0, left=None, right=None, next=None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next

class Solution(object):
    def connect(self, root):
        """
        :type root: Node
        :rtype: Node
        """
        if root is None:
            return root
        level = [root]
        while len(level) > 0:
            next = []
            for i in range(len(level)):
                n = level[i]
                if i < len(level)-1:
                    n.next = level[i+1]
                if n.left is not None:
                    next.append(n.left)
                if n.right is not None:
                    next.append(n.right)
            level = next
        return root

def walk(node):
    if node is None:
        return
    nv = None
    if node.next is not None:
        nv = node.next.val
    print(node.val, nv)
    walk(node.left)
    walk(node.right)

if __name__ == '__main__':
    sol = Solution()
    root = Node(1, Node(2, Node(4), Node(5)), Node(3, None, Node(7)))
    sol.connect(root)
    walk(root)
    print()
    root = Node(1, Node(2, Node(4)), Node(3, None, Node(5)))
    sol.connect(root)
    walk(root)
