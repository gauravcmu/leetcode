# Problem 
X is a good number if after rotating each digit individually by 180 degrees, we get a valid number that is different from X.  Each digit must be rotated - we cannot choose to leave it alone.

A number is valid if each digit remains a digit after rotation. 0, 1, and 8 rotate to themselves; 2 and 5 rotate to each other; 6 and 9 rotate to each other, and the rest of the numbers do not rotate to any other number and become invalid.

Now given a positive number N, how many numbers X from 1 to N are good?

Example:
Input: 10
Output: 4
Explanation: 
There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
Note that 1 and 10 are not good numbers, since they remain unchanged after rotating.

# Solution - 
#### 1. Bruteforce 

#### 2. Use math 
For a X digit number, we can say:-

total = number of X digit good numbers + number of X-1 digit good numbers .. so on. 
```
T = F(how many X digit number in N) +  F(X-1 digits) + F(x-2 digits)...
```

For X digits, the total good numbers can be found using permutations with the following formula:
```
F (x) = power(7, x) - pow(3, x)  
```

For example:- F(100)
F(100) = F'(100 - 100 is only 3 digit number in the input) + ( 7*7 - 3*3) + (7 - 3)  
F(100) = F'(100) + 40 + 4

F'(100) can be calculated in bruteforce way. 
