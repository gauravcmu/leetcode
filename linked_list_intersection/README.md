## Problem 

Write a program to find the node at which the intersection of two singly linked lists begins.


For example, the following two linked lists:
```
A:          a1 → a2
                   ↘
                     c1 → c2 → c3
                   ↗            
B:     b1 → b2 → b3
```
begin to intersect at node c1.


Notes:

If the two linked lists have no intersection at all, return null.
The linked lists must retain their original structure after the function returns.
You may assume there are no cycles anywhere in the entire linked structure.
Your code should preferably run in O(n) time and use only O(1) memory.

## Solution
	- Find lengths of the two lists - l1 and l2. 
	- find which list is bigger. Move l1-l2 nodes if A is longer list otherwise l2-l1 for list B. 
	- At this point, you are at two nodes, that are equidistant from the end node c3. 
	- go node by node on both lists and compare pointer. 
