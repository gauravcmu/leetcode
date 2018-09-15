# Problem
A peak element is an element that is greater than its neighbors.

Given an input array nums, where nums[i] ≠ nums[i+1], find a peak element and return its index.

The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.

You may imagine that nums[-1] = nums[n] = -∞.

Example 1:

Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
Example 2:

Input: nums = [1,2,1,3,5,6,4]
Output: 1 or 5 
Explanation: Your function can return either index number 1 where the peak element is 2, 
             or index number 5 where the peak element is 6.

### Solution

Use binary search to move towards a side based on next element to mid. 
If value at mid-1 is greater than mid, then search in left side else in right side. 
	if value at mid greater than its neighbors, thats the answer. 
	if start == end -> value at start is the answer. 
	if list is 2 numbers - end-start ==1 
		greater of the two is the answer. 

Explaination:
[1,2,1,3,5,6,4]

mid value = 3
3 less than 5 implies one of two things:
	1. 5 is greater than its neighbors. 
	2. else number next to 5 is greater than 5. 
	Either way, a peak element lies to right. Either it is 5 or its in the list beyond 5. 
	Thus explore right..

Take case where it keeps increasing in one direction:
[1,2,1,1,3,5,6,7,8]

Say mid value = 3.

Now, mid value is smaller than right neighbor - 5. So explore right side. 
findPeak(nums, start=6, end=8)
[5,6,7,8]

Now mid value = 6
again explore right. 
[7, 8]

greater of the two is 8, answer is 8. 

### proof that it works.
1. we're moving towards a side we see is greater than mid, while checking mid. 
2. The side being explored, appears to be on increasing trend. 
	- If the trend is maintained, i.e. all values after mid keep increasing - then end element is peak. 
	- if at any point the trend is broken, then the element before trend change is the peak. 

This proves that algorithm moves in the direction of the peak element. 

```
Write FindPeak(nums, start, end int) int {
	- if start == end  
		- return start element. 
	- if 2 element list, return bigger. 
		if end = star +1 
	- if mid is peak - check neighbors. 
	- if not, move towards bigger neighbor. 
}
```