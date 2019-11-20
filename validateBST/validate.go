/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true 
    }
    
    lvalid:= helper(root.Left, math.MinInt64, root.Val)
    rvalid:= helper(root.Right, root.Val, math.MaxInt64)
    
    return lvalid && rvalid 
}

func helper(root *TreeNode, min, max int) bool {
    if root == nil {
        return true 
    }
    
    //root left right - max min are passed down. 
    
    if root.Val <= min {
        return false 
    }
    
    if root.Val >= max {
        return false 
    }
    
    //my left subtree should be between min and max my value. 
    lvalid := helper(root.Left, min, root.Val) 
    //my right subtree should be between min my value and max. 
    rvalid := helper(root.Right, root.Val, max)
    
    return lvalid && rvalid 
}
