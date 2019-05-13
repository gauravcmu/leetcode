## Problem 

#### Binary Tree Level Order Traversal II
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7]

```
    3
   / \
  9  20
    /  \
   15   7
```
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]

### Solution 
Traverse level by level the tree. 
At each level, drain the q1 queue and insert nodes to q2. 
	once drained, drain q2 and for each element, insert it to array r - nodes at this level
	for each element in q2, insert its kids to original queue q1. 
Insert r to result set at first index. 

Return result. 
