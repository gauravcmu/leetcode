/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
	"fmt"
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	binaryTreePathsHelper(root, "", &result)
	return result
}

func binaryTreePathsHelper(root *TreeNode, path string, result *[]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*result = append(*result, path+strconv.Itoa(root.Val))
		fmt.Printf("result:%+v\n", result)
	} else {
		if root.Left != nil {
			binaryTreePathsHelper(root.Left, path+strconv.Itoa(root.Val)+"->", result)
		}
		if root.Right != nil {
			binaryTreePathsHelper(root.Right, path+strconv.Itoa(root.Val)+"->", result)
		}
	}
}