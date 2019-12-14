package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	q1 := &queue{}
	q2 := &queue{}

	if root == nil {
		return [][]int{}
		
	}

	q1.enqueue(root)

	lr := false
	result := make([][]int, 0)
	for q1.isEmpty() != true {
		lr = !lr
		for q1.isEmpty() != true {
			q2.enqueue(q1.dequeue())
		}

		res := make([]int, 0)
		for q2.isEmpty() != true {
			v := q2.dequeue()
			q1.enqueue(v.Left)
			q1.enqueue(v.Right)
			if lr {
				res = append(res, v.Val)
			} else {
				//reverse array every other time
				res = append([]int{v.Val}, res...)
			}
		}
		result = append(result, res)
	}
	return result
}

type queue []*TreeNode

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *queue) enqueue(node *TreeNode) {
	if node != nil {
		*q = append(*q, node)
	}
}

func (q *queue) dequeue() *TreeNode {
	t := (*q)[0]
	*q = (*q)[1:]
	return t
}
