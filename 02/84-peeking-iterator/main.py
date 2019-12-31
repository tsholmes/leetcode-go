# Below is the interface for Iterator, which is already defined for you.
#
class Iterator(object):
    def __init__(self, nums):
        """
        Initializes an iterator object to the beginning of a list.
        :type nums: List[int]
        """
        self.nums = nums

    def hasNext(self):
        """
        Returns true if the iteration has more elements.
        :rtype: bool
        """
        return len(self.nums) > 0

    def next(self):
        """
        Returns the next element in the iteration.
        :rtype: int
        """
        v = self.nums[0]
        self.nums = self.nums[1:]
        return v

class PeekingIterator(object):
    def __init__(self, iterator):
        """
        Initialize your data structure here.
        :type iterator: Iterator
        """
        self.it = iterator
        self.p = None

    def peek(self):
        """
        Returns the next element in the iteration without advancing the iterator.
        :rtype: int
        """
        if self.p is None:
            self.p = self.it.next()
        return self.p

    def next(self):
        """
        :rtype: int
        """
        if self.p is not None:
            v = self.p
            self.p = None
            return v
        return self.it.next()

    def hasNext(self):
        """
        :rtype: bool
        """
        return self.p is not None or self.it.hasNext()


# Your PeekingIterator object will be instantiated and called as such:
# iter = PeekingIterator(Iterator(nums))
# while iter.hasNext():
#     val = iter.peek()   # Get the next element but not advance the iterator.
#     iter.next()         # Should return the same value as [val].

if __name__ == '__main__':
    it = Iterator([1,2,3,4,5])
    pit = PeekingIterator(it)
    while pit.hasNext():
        print(pit.peek())
        print(pit.next())
