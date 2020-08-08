class StringIterator(object):

    def __init__(self, compressedString):
        """
        :type compressedString: str
        """
        self.pairs = []

        i = 0
        while i < len(compressedString):
            letter = compressedString[i]
            num = 0
            i = i + 1
            while i < len(compressedString) and compressedString[i] >= '0' and compressedString[i] <= '9':
                num = num * 10
                num = num + int(compressedString[i])
                i += 1
            self.pairs.append([letter, num])

    def next(self):
        """
        :rtype: str
        """
        if len(self.pairs) == 0:
            return ' '
        l = self.pairs[0][0]
        if self.pairs[0][1] == 1:
            self.pairs.pop(0)
        else:
            self.pairs[0][1] -= 1
        return l

    def hasNext(self):
        """
        :rtype: bool
        """
        return len(self.pairs) > 0



# Your StringIterator object will be instantiated and called as such:
# obj = StringIterator(compressedString)
# param_1 = obj.next()
# param_2 = obj.hasNext()
