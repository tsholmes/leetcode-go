"""
# Definition for a Node.
class Node(object):
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""
"""
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
"""
class Codec:

    def encode(self, root):
        """Encodes an n-ary tree to a binary tree.
        :type root: Node
        :rtype: TreeNode
        """
        if root is None:
            return None

        def build(n):
            res = TreeNode(n.val)
            r = None
            for c in (n.children or []):
                n2 = build(c)
                if r is None:
                    res.right, r = n2, n2
                else:
                    r.left = n2
                    r = n2
            return res
        return build(root)

    def decode(self, data):
        """Decodes your binary tree to an n-ary tree.
        :type data: TreeNode
        :rtype: Node
        """
        if data is None:
            return None

        def build(n):
            res = []
            while n is not None:
                rn = Node(n.val)
                rn.children = build(n.right)
                n = n.left
                res.append(rn)
            return res

        return build(data)[0]


# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.decode(codec.encode(root))
