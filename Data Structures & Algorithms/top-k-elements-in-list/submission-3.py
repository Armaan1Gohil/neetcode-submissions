class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        counter = {}
        for num in nums:
            counter[num]  = 1 + counter.get(num, 0)
        
        freq = [[] for i in range(len(nums) + 1)]
        for n, c in counter.items():
            freq[c].append(n)
        
        res = []
        while k:
            last_ele = freq.pop()
            if last_ele:
                for n in last_ele:
                    res.append(n)
                k -= len(last_ele)
        return res

