# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.

        :type root: TreeNode
        :rtype: str
        """
        res = []
        def walk(n):
            if n is None:
                res.append('')
                return
            res.append(str(n.val))
            walk(n.left)
            walk(n.right)
        walk(root)
        return ','.join(res)


    def deserialize(self, data):
        """Decodes your encoded data to tree.

        :type data: str
        :rtype: TreeNode
        """
        vals = data.split(',')
        i = [0]
        def build():
            j = i[0]
            i[0] += 1
            if vals[j] == '':
                return None
            n = TreeNode(int(vals[j]))
            n.left = build()
            n.right = build()
            return n
        return build()


# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))

if __name__ == '__main__':
    n = TreeNode(2)
    n.left = TreeNode(1)
    n.right = TreeNode(3)
    c = Codec()
    print(c.serialize(n))
    print(c.deserialize(c.serialize(n)))
