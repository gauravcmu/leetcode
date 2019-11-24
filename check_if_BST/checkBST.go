package main

import (
	"fmt"
)

/* 
 A tree is BST if the value at a node is greater than all values in its left subtree. 
 A tree is BST if value at the node is smaller than all the values in its right subtree. 
*/

func verifyBST(root *node, lowest, highest int) bool {
 if root == nil {
   return true 
 }

   isLeftBST := verify(root.Left, lowest, root.Val)  
   isRightBST := verify(root.Right, root.Val, highest)
  
   return isLeftBST && isRightBST && (root.Val > lowest && root.Val < highest)
}
