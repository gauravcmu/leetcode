## Problem 

Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]

## Solution

1. Use a queue for level order traversal. Use a stack for handling the queue in LIFO order.
2. Use a variable to track left to right vs right to left direction - use this while adding children to queue. 

	- Enqueue the root. 
	- Initialize LR to true. 
	- empty the queue (i.e. entire level)
		- Dequeue v from queue
			- push v to stack. 
		- Pop the entire stack now. 
			- for each v popped from stack. 
				- insert its children left to right or right to left based on LR. 
	- Once a level is processed - toggle LR flag for next level. 
	