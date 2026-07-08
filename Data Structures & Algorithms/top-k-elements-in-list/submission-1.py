class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        counter = {}
        for num in nums:
            counter[num]  = 1 + counter.get(num, 0)
        
        sorted_counter = dict(sorted(counter.items(), key=lambda key_val: key_val[1], reverse=True))
        
        res = []
        for i in range(k):
            res.append(list(sorted_counter.keys())[i])
        
        return res
            
