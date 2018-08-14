# Problem

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
```
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```
## Solution:
- Sort the array. 
```
    [ -2 -2 -1 0 1 2 3 4 5 ]
       ^  ^ --->    <--- ^
       i  s              e
```
- maintain 3 pointers. 
	- `i: 0 to len-2 `
	- `s: i+1 towards end (s < e)`
	- `e: len-1 towards 0`
- add to result when values at i, s and e add up to 0. 
	- Else. If sum is less than 0 > move start towards right. 
	- Sum is greater than 0, then move end towards left. 
	- while moving, avoid same element as current value..

