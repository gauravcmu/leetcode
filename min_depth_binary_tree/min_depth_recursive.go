/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return helper(root)
}


/*
    nil root - return max value as we do not want this to be calculated in min from the caller. 
    if leaf node - than return 1. 
    else - return min(left / right subtree depth) + 1
*/
func helper(root * TreeNode) int {
    if root == nil {
        return math.MaxInt32 
    }
    
    if root.Left == nil && root.Right == nil {
        return 1
    }
    
    lmin := helper(root.Left)
    rmin := helper(root.Right)
    
    return min(lmin, rmin) + 1
}

func min(a, b int) int {
    if a < b {
        return a 
    }
    return b 
}

/*
helper(3)
    helper(9)
        returns 1 
    helper(20) - return 2 
        helper(15)
            returns 1 
        helper(7)
            returns 1   
*/
