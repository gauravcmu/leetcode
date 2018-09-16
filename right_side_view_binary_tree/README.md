## Problem

Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

Example:

Input: [1,2,3,null,5,null,4]
Output: [1, 3, 4]
Explanation:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---


 	1
  /   \             = Solution should be 1,3,5. 
 2		3
   \
 	5


### Solution

Right side is right most node in level order traversal. 

At each level, find the last node in queue and insert this in result set. 
To find the last node at each level -
	- in the BFS loop - while queue-1 not empty
		- each iteration, drain the queue completely (entire level)
			- cache last from the level using last() - this is rightmost at this level.
			- insert this level node by node in a separate queue - q2. 
		- go over q2 until empty
			- for each v in q2, insert its children left to right in q1.
			- ^ above builds next level. 
return result. 