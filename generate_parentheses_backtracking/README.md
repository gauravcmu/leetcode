# Problem

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

# Solution approach

Use backtracking. 
1. if in any tree path, open is greater than closed, it is valid. 
2. If closed is greater than open, it is invalid path. Do not explore. 
3. If open == closed == n (number of pairs) - add the output to result

https://www.youtube.com/watch?v=pD2VNU2x8w8