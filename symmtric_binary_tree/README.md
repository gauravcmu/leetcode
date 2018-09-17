### Problem 
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
```
    1
   / \
  2   2
 / \ / \
3  4 4  3
But the following [1,2,2,null,3,null,3] is not:
    1
   / \
  2   2
   \   \
   3    3
```
### Solution

Write a function isMirror(root1, root2) bool
- return true if both trees null
- if root1 and root2 are non-nil 
	- if root1.Val == root2.Val
		- return isMirrot(r1.Left, r2.Right) && isMirror(r1.Right, r2.Left) 
return false. 
