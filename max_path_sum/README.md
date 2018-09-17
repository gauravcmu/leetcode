### Problem
Given a binary tree, find the maximum path sum. The path may start and end at any node in the tree.

Example:

Input: Root of below tree
```
       1
      / \
     2   3
```
Output: 6

See below diagram for another example.
1+2+3
tree

### Solution. 

Consider 4 cases for a path going through a node:
1. path stops at the node. 
2. path goes throught node to left subtree.
3. path goes throught node to right subtree. 
4. path comes from left / right subtree - goes to node and then to right / left subtree. 

Max of 1, (2,3) and 4 -> max_top. Now maximize result with max_top but function returns max_single (1,2,3). 
Return incorporates 1,2, and 3 only so that the parent can use this path. 4 is invalid if parent uses this path. 

  
