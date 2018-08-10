class Solution(object):
    def countPrimes(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n <= 2: 
            return 0 #1 is special, 2 is prime but less than 2 are 0 
        if n ==3:
            return 1 # only 2 is prime 
        
        count = 2 #below 5, only 2 and 3 are primes - count from 2 
        i = 5 
        
        while (i < n):
            if self.isPrime(i) == True:
                count+=1 
            i+=2   # just the odd numbers. Why check even numbers.. 
        return count 
    
    def isPrime(self, n):
        """return true if prime else false"""
        #first 3 checks are redundant due to caller handling them though. 
        if n == 0: 
            return False
        if n == 1: 
            return False
        if n <=3:
            return True
        
        if n%2 == 0 or n%3==0:
            return False
        #number does not have a factor of the form 6k-1 or 6k+1
        j=5
        while (j*j<=n):
            if (n%j == 0) or (n%(j+2) == 0):
                return False 
            j+=6
        return True
