# Problem
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:
```
Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
```
## Solution

Create a stack of strings. 
Create a 2d array of pairs of opening and closing parenthesis.
``` 
var paran [][]string = [][]string { {"(", ")"}, {"{", "}"}, {"[", "]"}} 
```
To validate - 
- Go over char by char in the string. 
	- check if char is opening paranthesis. If so, push it on the stack. 
	- if it is closing parenthesis - find an entry in array that has opening and closing in same entry in the 2d array. 
	- if found a match, pop and move on to next. 
		- if end of input, then stack should be empty for valid paranthesis. 

