/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import ("strconv")
func binaryTreePaths(root *TreeNode) []string {
    
    if root == nil {
        return []string{}
    }
    
    left := binaryTreePaths(root.Left)
    right := binaryTreePaths(root.Right)
    var paths []string 
    
    if root.Left != nil {
        for _, val := range left {
            paths = append(paths, strconv.Itoa(root.Val) + "->" + val) 
        }
    }    
    
    if root.Right != nil {
        for _, val := range right {
            paths = append(paths, strconv.Itoa(root.Val) + "->" + val) 
        }
    }    

    if root.Left == nil && root.Right == nil {
        paths = []string{strconv.Itoa(root.Val)}
    }
    return paths 

}
