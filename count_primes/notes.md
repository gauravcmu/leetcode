Count the number of prime numbers less than a non-negative number, n.

Example:

Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.

Approach to solve:
- Treat 0, 1, 2, 3 as special cases, return 0 and 1 respectively.
- From 5 onwards - check odd numbers for primality. 
Primality check:
- Treat 0-3 with if checks. (actually redundant due to caller checking)
- Check if divisible by 2 or 3 
- From j = 5 (6k-1) to j*j <=n - check if j is a factor. 
	- for 6k-1 => increment j by 6
	- for 6k+1 - use j+2 as j already 6k-1.   

Mistakes - 
- treating 1 as prime
- 2 and 3 are primes. 
- forgetting check for divisibility by 2 and by 3. 
- 
optimize - 
- can loop over only odd numbers.  

Tests that failed:
Input:
10000
Output:
1241
Expected:
1229

25, 121, 289, 529, 841, 1681, 2209, 2809, 3481, 5041, 6889, 7921 
- ^ treated as prime - bug was checking j*j < n instead of j*j<=n 
