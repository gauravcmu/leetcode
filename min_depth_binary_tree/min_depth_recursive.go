/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 
    1 + min(1, 2)
    md(3)
        1 ld- md(9)
           0 - md(nil)
        2- rd - md(20) 
            1 - md(15)
               0 - md(nil)
            1- md(7)   
               0 - md(nil)
 */
func minDepth(root *TreeNode) int {
    // root min left / min right
    if root == nil {
        return 0 
    } 
    
    //post order traversal. left right root.
    ld := minDepth(root.Left)
    rd := minDepth(root.Right)
    
    // 1 + min (ld, rd)
    if ld !=0 && rd != 0 {
        return min(ld, rd) + 1
    }
        return ld + rd + 1
}

func min (a, b int) int {
    if a < b {
        return a
    }
    return b
}
