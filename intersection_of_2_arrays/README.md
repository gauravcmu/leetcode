#Problem

Given two arrays, write a function to compute their intersection.

Example 1:
```
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
```
Note:
Each element in the result must be unique.
The result can be in any order.
 
## Solution:
	- for each entry in nums1 - create hash table entry with value 1
	- iterate over nums2 - for each entry already in the hash, mark as -1 value
	- iterate over the mapping and look for -1 - put this in result array. 
