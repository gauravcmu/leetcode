import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	var q *queue = &queue{}
	var q2 *queue = &queue{}
	q.Enqueue(root)

	for !q.IsEmpty() {
		last := q.Last()
		for !q.IsEmpty() {
			v := q.Dequeue()
			q2.Enqueue(v)
		}

		for !q2.IsEmpty() {
			v := q2.Dequeue()
			q.Enqueue(v.Left)
			q.Enqueue(v.Right)
		}
		//this is right most node.
		res = append(res, last.Val)
	}
	return res
}

type queue []*TreeNode

func (q *queue) Last() *TreeNode {
	last := (*q)[len(*q)-1]
	fmt.Printf("last:%v\n", last)
	return last
}
func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue) Enqueue(value *TreeNode) {
	if value == nil {
		return
	}
	*q = append(*q, value)
	fmt.Printf("enqueue:%+v queue:%+v\n", value, *q)
}

func (q *queue) Dequeue() *TreeNode {
	t := (*q)[0]
	(*q) = (*q)[1:]
	return t
}