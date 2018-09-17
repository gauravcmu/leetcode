import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	return hasPathSumHelper(root, sum)
}

func hasPathSumHelper(root *TreeNode, sum int) bool {
	fmt.Printf("root.Val:%v, sum:%v\n", root, sum)
	if root == nil {
		return false
	}

	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}

	lflag := hasPathSumHelper(root.Left, sum-root.Val)
	rflag := hasPathSumHelper(root.Right, sum-root.Val)

	return lflag == true || rflag == true
}