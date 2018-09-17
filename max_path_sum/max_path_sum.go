/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"

func maxPathSum(root *TreeNode) int {
	result := math.MinInt32
	maxPathSumHelper(root, &result)
	return result
}
func maxPathSumHelper(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	lmax := maxPathSumHelper(root.Left, result)
	rmax := maxPathSumHelper(root.Right, result)

	//max of:
	//root.Val (no child in path)
	//root.Val + subtree with max path.
	max_atmost_single_child := max(max(lmax, rmax)+root.Val, root.Val)

	//max with left right and root.
	//or max with at most one single.
	max_top := max(max_atmost_single_child, lmax+rmax+root.Val)

	*result = max(*result, max_top)

	//root can use this only when one path is taken in subtree. Thus path with
	//single child is used.
	return max_atmost_single_child
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}