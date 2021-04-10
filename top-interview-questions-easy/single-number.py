class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        dict = defaultdict(int)
        
        for i in nums:
            dict[i] += 1

        for i in dict:
            if dict[i] == 1:
                return i
