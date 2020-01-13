# The guess API is already defined for you.
# @param num, your guess
# @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
# def guess(num):

class Solution(object):
    def guessNumber(self, n):
        """
        :type n: int
        :rtype: int
        """
        low, high = 1, n
        while low < high:
            mid = (low + high) / 2
            g = guess(mid)
            if g == -1:
                high = mid - 1
            elif g == 1:
                low = mid + 1
            else:
                low, high = mid, mid
        return low
