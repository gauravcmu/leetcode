/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return diaTree(root) -1 
}

func diaTree(root *TreeNode) int {
    if root == nil {
        return 0 
    }
    
    lh := height(root.Left)
    rh := height(root.Right)
    
    l := diaTree(root.Left)
    r := diaTree(root.Right)
   
    return max(lh+rh+1, max(l, r))  
}

func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(height(root.Left), height(root.Right))  
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
