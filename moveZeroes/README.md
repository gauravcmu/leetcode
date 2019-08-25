Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.

## Solution
===========

Iterate over the array and maintain index to write at:
	- Now write non-zero element at index_to_write_at
		- increment index_to_write_at
	- just move to next element if it is zero. 
Once you have done this loop - loop from index_to_write_at to end of the array
	- set each element to zero. 

