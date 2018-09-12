## Problem

Two sum - get indices of two numbers in integer array that sum up to a target value. 
-	Cannot use the same number twice. 

### Solution 
1. Use a hash table to map numbers as keys 
	- Lookup in hash table (target - value) and if found, return index.
2. Use an array with original index an value stored. 
	- Sort the array. 
	- Keep two pointers, low and high. 
		- if nums[low] + nums[high] == target - return
		- if ^^^^ < target - low++
		- else high-- 
	O(log n) for sort + O(n) for compare and find step. O(n) extra space.
