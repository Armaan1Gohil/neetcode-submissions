class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        check = set()
        res = 0
        l = 0
        for r in range(len(s)):
            while s[r] in check:
                check.remove(s[l])
                l += 1
            check.add(s[r])
            res = max(res, r - l + 1)
        return res