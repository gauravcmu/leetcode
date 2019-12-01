/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"
func maxPathSum(root *TreeNode) int {
    t := math.MinInt32
    helper(root, &t) 
    return t 
}

/*
Return what value will it be if we add a path going via this root and one of its subtrees. 
While doing this, maximize result, which is max of result and root+left+right. 
*/
func helper(root * TreeNode, result *int) int {
    if root == nil {
        return 0 
    }
    
    lmax := helper(root.Left, result) 
    rmax := helper(root.Right, result)
    
    //So at this node, we could have max value or path sum to be equal to
    //sum of this node + left and right. 
    *result = max(*result, lmax + rmax + root.Val)
    
    //But if we contribute to a path sum at above levels, we could only use 
    //either left or right subtree path. Ensure we dont go negative
    return max(0, root.Val + max(lmax, rmax))
}

func max(a, b int) int {
    if a > b {
        return a
    } 
    return b 
}
