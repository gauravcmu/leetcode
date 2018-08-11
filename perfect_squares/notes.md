## Problem Statement
Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

Example 1:

```Input: n = 12
Output: 3 
Explanation: 12 = 4 + 4 + 4.
Example 2:

Input: n = 13
Output: 2
Explanation: 13 = 4 + 9. 
```

## Solution
- Find the biggest perfect square below N. 
	- Do sqrt(n) - should give sqrt (floor value) 
- N = biggest perfect square (1) + F(N- bps) 
- So, iterate from bps root to 1 and minimize sum. 

For example:- 
12 - bps 9 and thus we need to minimize with 3 -> 1
```
F(12) = 1 + F(N - (3...1))
       MIN( 1 + F(12 - 3*3) OR 1 + F(12 - 2*2) OR 1 + F(12-1*1))
```
	
