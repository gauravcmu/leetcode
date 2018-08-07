class Solution(object):
    def __init__(self): 
        self.map = {"2":"abc", "3":"def", "4":"ghi", "5":"jkl", "6":"mno", "7":"pqrs", "8":"tuv", "9":"wxyz"}
    def letterCombinations(self, digits):
        """
        :type digits: str
        :rtype: List[str]
        """
        result = []
        l = len(digits)
        print("digits: len:", digits, l)
        if l == 0:
            return result 
        if l == 1: 
            s = self.map[digits]
            for v in s: 
                result.append(v)    
            return result       
        mid = l/2
        left = self.letterCombinations(digits[:l/2]) 
        print(left)
        right = self.letterCombinations(digits[(l/2):])
        print(right)
        for v1 in left:
            for v2 in right: 
                result.append(v1+v2) 
                print("---result--", result)
                
        print("result:",result)
        return result         
