package lowest_common_ancestor

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return helper(root, p, q)
}

func helper(root, p, q *TreeNode) *TreeNode {
	//empty tree. get back.
	if root == nil {
		return nil
	}

	//see if left has a match or right has a match.
	l := helper(root.Left, p, q)
	r := helper(root.Right, p, q)

	//left and right both have one of the values each. assumes no duplicates..
	if l != nil && r != nil {
		return root
	}

	//if root has any values, then return root.
	if root == p || root == q {
		return root
	}

	// return the node that is set.
	if l == nil {
		return r
	}
	return l
}
