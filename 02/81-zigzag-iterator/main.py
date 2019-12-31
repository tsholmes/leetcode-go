class ZigzagIterator(object):

    def __init__(self, v1, v2):
        """
        Initialize your data structure here.
        :type v1: List[int]
        :type v2: List[int]
        """
        self.vs = [v1, v2]
        self.pos = 0 if len(v1) > 0 else 1
        self.i = 0

    def next(self):
        """
        :rtype: int
        """
        res = self.vs[self.pos][self.i]
        if len(self.vs[1-self.pos]) <= self.i + self.pos:
            self.i += 1
        else:
            self.i += self.pos
            self.pos = 1 - self.pos
        return res

    def hasNext(self):
        """
        :rtype: bool
        """
        return self.i < len(self.vs[self.pos])

# Your ZigzagIterator object will be instantiated and called as such:
# i, v = ZigzagIterator(v1, v2), []
# while i.hasNext(): v.append(i.next())
