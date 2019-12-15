# There is no go submissions for this problem :(
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
        depth = 0
        left = root
        while left is not None:
            left = left.left
            depth += 1
        def get(i):
            node = root
            d = depth
            while (1<<d) > i:
                d -= 1
            for b in range(d-1, -1, -1):
                if (i&(1<<b)) == 0:
                    node = node.left
                else:
                    node = node.right
            return node
        for i in range((1<<depth) - 1):
            x = i + 1
            y = i + 2
            if (y & -y) != y:
                get(x).next = get(y)
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
    root = Node(1, Node(2, Node(4), Node(5)), Node(3, Node(6), Node(7)))
    sol.connect(root)
    walk(root)
