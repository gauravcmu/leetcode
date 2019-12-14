package isBalancedBinaryTree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	tb, _ := isBal(root)
	return tb
}

func isBal(root *TreeNode)  (bool, int){
	h := -1
	if root == nil {
		return true, 0
	}

	lb, lh := isBal(root.Left)
	rb, rh := isBal(root.Right)

	if lb&&rb && ((lh - rh)*(lh - rh) <= 1) {
		h = max(lh, rh) + 1
		return true, h
	} else {
		return false, h
	}
}

func max (a, b int) int {
	if a > b {
		return a
	}
	return b
}
