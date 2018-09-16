import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var q *queue = &queue{}
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	q.Enqueue(root)

	var q2 *queue = &queue{}

	for !q.IsEmpty() {
		for !q.IsEmpty() {
			v := q.Dequeue()
			q2.Enqueue(v)
		}
		r := make([]int, 0)
		for !q2.IsEmpty() {
			v := q2.Dequeue()
			r = append(r, v.Val)
			q.Enqueue(v.Left)
			q.Enqueue(v.Right)
		}
		fmt.Printf("r:%+v\n", r)
		res = append([][]int{r}, res...)
	}
	return res
}

type queue []*TreeNode

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue) Enqueue(value *TreeNode) {
	if value == nil {
		return
	}
	*q = append((*q), value)
}

func (q *queue) Dequeue() *TreeNode {
	t := (*q)[0]
	(*q) = (*q)[1:]
	return t
}

