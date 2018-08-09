#Given a collection of distinct integers, return all possible permutations.
#
#Example:
#
#Input: [1,2,3]
#Output:
#[
#  [1,2,3],
#  [1,3,2],
#  [2,1,3],
#  [2,3,1],
#  [3,1,2],
#  [3,2,1]
#]
# Solution: use backtracking. Chose - explore - unchose. 

class Solution:
    def __init__(self):
        self.result = []
    def permute(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        c = self.permuteHelper(nums, [])
        return self.result 
       
    def permuteHelper(self, nums, chosen):
        if len(nums) == 0:
            self.result.append(copy.copy(chosen))
            return self.result 
        index = 0 
        for n in nums:
            nums.pop(index)
            chosen.append(n)
            
            c = self.permuteHelper(nums, chosen)
            
            nums.insert(index, n)
            chosen.remove(n)
            index+=1
