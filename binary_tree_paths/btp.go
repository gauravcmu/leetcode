/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
    result := make([]string, 0)
    if root == nil {
        return result 
    }
    
    if root != nil && root.Left == nil  && root.Right == nil {
		result = append(result, strconv.Itoa(root.Val))
		return result
	}
    paths(root.Left, strconv.Itoa(root.Val), &result)
    paths(root.Right, strconv.Itoa(root.Val), &result)
    return result 
}

/*
paths(1, "")
    paths(2, )
*/
func paths(root *TreeNode, path string, result *[]string) {
    fmt.Printf("result:%v\n", result)
    if root == nil {
        return
    }
	if root != nil && root.Left == nil  && root.Right == nil {
		*result = append(*result, path+"->"+strconv.Itoa(root.Val))
		return
	}
    paths(root.Left, path+"->"+strconv.Itoa(root.Val), result )
    paths(root.Right, path+"->"+strconv.Itoa(root.Val), result)
}
