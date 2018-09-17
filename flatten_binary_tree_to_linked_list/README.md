### Problem

Flatten Binary Tree to Linked List

Given a binary tree, flatten it to a linked list in-place.

For example, given the following tree:
```
    1
   / \
  2   5
 / \   \
3   4   6
```
The flattened tree should look like:
```
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
```

### Solution. 

Imaging a simple tree as shown:
```  
    1 <-root
  l/ \r
  2   5
  ```

To   
```
1 --Right--> 2 --Right--> 5


Root -> Right = (flattened left subtree) --Right--> (flattened right subtree)
root -> left = nil 
```