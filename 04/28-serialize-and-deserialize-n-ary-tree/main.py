# Definition for a Node.
class Node(object):
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class Codec:
    def append_varint(self, res, n):
        n = n & 0xFFFFFFFF
        if n >= 1<<24:
            res.append(chr(4))
            res.append(chr((n>>24)&0xFF))
            res.append(chr((n>>16)&0xFF))
            res.append(chr((n>>8)&0xFF))
            res.append(chr((n>>0)&0xFF))
        elif n >= 1<<16:
            res.append(chr(3))
            res.append(chr((n>>16)&0xFF))
            res.append(chr((n>>8)&0xFF))
            res.append(chr((n>>0)&0xFF))
        elif n >= 1<<8:
            res.append(chr(2))
            res.append(chr((n>>8)&0xFF))
            res.append(chr((n>>0)&0xFF))
        else:
            res.append(chr(1))
            res.append(chr((n>>0)&0xFF))

    def read_varint(self, data, i):
        ilen = ord(data[i])
        i += 1
        res = 0
        for _ in range(ilen):
            res = res << 8
            res = res | ord(data[i])
            i += 1
        if res & 1 << 31 != 0:
            res = (~res & 0x7FFFFFFF) + 1
            res = -res
        return res, i

    def serialize(self, root):
        """Encodes a tree to a single string.

        :type root: Node
        :rtype: str
        """
        if root is None:
            return ""

        def walk(n, res):
            self.append_varint(res, n.val)
            self.append_varint(res, len(n.children or []))
            for i in range(len(n.children or [])):
                walk(n.children[i], res)

        res = []
        walk(root, res)
        return ''.join(res)

    def deserialize(self, data):
        """Decodes your encoded data to tree.

        :type data: str
        :rtype: Node
        """
        if data == "":
            return None

        def walk(i):
            val, i = self.read_varint(data, i)
            cc, i = self.read_varint(data, i)
            cs = []
            for _ in range(cc):
                c, i = walk(i)
                cs.append(c)
            return Node(val, cs), i

        n, _ = walk(0)
        return n


# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))

if __name__ == '__main__':
    codec = Codec()
    ser = codec.serialize(Node(1, [Node(2), Node(3), Node(4)]))
    print([ord(c) for c in ser])
    print(codec.deserialize(ser).children[1].val)
