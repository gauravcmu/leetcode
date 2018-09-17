/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	temp := root.Left
	for temp != nil && temp.Right != nil {
		temp = temp.Right
	}

	if temp != nil {
		temp.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
}
