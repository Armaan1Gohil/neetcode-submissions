class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        counter = {} 
        for num in nums:
            counter[num] = 1 + counter.get(num, 0)
        
        for key in counter:
            if counter[key] > 1:
                return True
        return False