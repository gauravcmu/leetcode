/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    } 
    if root.Left == nil && root.Right == nil {
        //leaf 
        return 1
    } 
    l := maxDepth(root.Left)
    r := maxDepth(root.Right)
    depth := max(l, r)
    return depth+1
}

func max(a,b int) int {
    if a > b {
        return a
    } 
    return b  
}
