class Solution:
    def trap(self, height: List[int]) -> int:
        length = len(height)

        max_left = [0] * length
        max_right = [0] * length

        for l in range(1, length):
            max_left[l] = max(height[l - 1], max_left[l - 1])
        for r in range(length - 2, -1, -1):
            max_right[r] = max(height[r + 1], max_right[r + 1])

        res = 0
        for i in range(length):
            min_length = min(max_left[i], max_right[i])
            if min_length - height[i] > 0:
                res += min_length - height[i]
        return res