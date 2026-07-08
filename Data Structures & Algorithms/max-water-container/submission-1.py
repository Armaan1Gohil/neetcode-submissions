class Solution:
    def maxArea(self, heights: List[int]) -> int:
        res = 0
        l, r = 0, len(heights) - 1

        while l < r:
            height = min(heights[l], heights[r])
            breath = r - l
            area = height * breath
            res = max(res, area)
            if heights[l] > heights[r]:
                r -= 1
            else:
                l += 1
        return res