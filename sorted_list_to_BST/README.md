### Problem -  Convert Sorted List to Binary Search Tree

Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted linked list: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:
```
      0
     / \
   -3   9
   /   /
 -10  5
```

#### Solution

Recursively build binary search tree. 
Since the list is sorted, from the mid point of the list as root, anything on the left should belong to left subtree and similarly, anything on the right to the right subtree. 

1. Count number of nodes in original sorted list. 
	- While counting, insert this into a hash table for quick lookup. 
		```[node_count]*ListNode```
2. Call helper Helper(head, 0, count, mapping)
3. helper()
	- if head is nil, return nil. 
	- if start == end (single node list)
		make a tree node and set left right as nil and return. 
	- if start < end 
		find mid by doing start + (end-start)/2
		get midNode from mapping. 
		build newNode.
		newNode.Left = Recurse Helper(head, start, mid-1, mapping)

		newNode.Right = Recurse Helper(midnode->next, mid+1, end, mapping)	
		return newNode. 
