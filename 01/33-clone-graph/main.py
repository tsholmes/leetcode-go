# no go submissions :(
# Definition for a Node.
class Node(object):
    def __init__(self, val, neighbors):
        self.val = val
        self.neighbors = neighbors

class Solution(object):
    def cloneGraph(self, node):
        """
        :type node: Node
        :rtype: Node
        """
        mapfwd = {}
        mapback = {}
        def walk(n):
            nid = id(n)
            if nid in mapfwd:
                return
            nn = Node(n.val, [])
            mapfwd[nid] = nn
            mapback[id(nn)] = n
            for n2 in n.neighbors:
                walk(n2)
        walk(node)
        for _, v in mapfwd.items():
            v.neighbors = [mapfwd[id(n)] for n in mapback[id(v)].neighbors]
        return mapfwd[id(node)]

if __name__ == '__main__':
    n1, n2, n3, n4 = Node(1, []), Node(2, []), Node(3, []), Node(4, [])
    n1.neighbors = [n2, n4]
    n2.neighbors = [n1, n3]
    n3.neighbors = [n2, n4]
    n4.neighbors = [n3, n1]

    sol = Solution()
    print(sol.cloneGraph(n1))
