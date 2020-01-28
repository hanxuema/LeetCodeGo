class Solution(object):
    def firstBadVersion(self, n):
        if n == 0:
            return 1
        left = 0
        right = n
        while left < right:
            mid = left + (right - left) // 2
            if isBadVersion(mid):
                right = mid
            else:
                left = mid + 1
            
        return left

  
    def firstBadVersion2(self, n):
        if n == 0:
            return 1
        left = 0
        right = n
        while left <= right:
            mid = left + (right - left) // 2
            if isBadVersion(mid):
                right = mid -1
            else:
                left = mid + 1
            
        return left