import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	var q *queue = &queue{}
	var s *stack = &stack{}

	if root == nil {
		return res
	}

	q.Enqueue(root)

	var lr bool = true

	for q.IsEmpty() != true {
		for q.IsEmpty() != true {
			v := q.Dequeue()
			s.Push(v)
		}

		r := make([]int, 0)
		for s.IsEmpty() != true {
			v := s.Pop()
			r = append(r, v.Val)
			if lr == true {
				fmt.Printf("Left to Right\n")
				q.Enqueue(v.Left)
				q.Enqueue(v.Right)
			} else {
				fmt.Printf("Right to left\n")
				q.Enqueue(v.Right)
				q.Enqueue(v.Left)
			}
		}
		lr = !lr
		if len(r) != 0 {
			fmt.Printf("Appending %v to %v\n", r, res)
			res = append(res, r)
		}
	}
	return res
}

type queue []*TreeNode

func (q *queue) Enqueue(value *TreeNode) {
	if value == nil {
		return
	}
	(*q) = append((*q), value)
}

func (q *queue) IsEmpty() bool {
	fmt.Printf("Len of queue:%v\n", len((*q)))
	return len((*q)) == 0
}

func (q *queue) Dequeue() *TreeNode {
	t := (*q)[0]
	(*q) = (*q)[1:]
	return t
}

type stack []*TreeNode

func (s *stack) Push(value *TreeNode) {
	if value == nil {
		return
	}
	fmt.Printf("push to stack %v\n", value)
	(*s) = append((*s), value)
}

func (s *stack) Pop() *TreeNode {
	t := (*s)[len((*s))-1]
	(*s) = (*s)[:len((*s))-1]
	return t
}

func (s *stack) Peek() *TreeNode {
	return (*s)[len((*s))-1]
}

func (s *stack) IsEmpty() bool {
	return len((*s)) == 0
}