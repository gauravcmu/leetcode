/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 != nil && root2 != nil && root1.Val == root2.Val {
		return isMirror(root1.Left, root2.Right) && isMirror(root2.Left, root1.Right)
	}
	return false
}