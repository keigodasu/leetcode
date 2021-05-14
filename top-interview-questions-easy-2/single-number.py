class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        dic = dict()
        
        for i in nums:
            dic[i] = dic.get(i, 0) + 1
            
        for num in dic:
            if dic[num] == 1:
                return num

