/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    
    if root == nil {
        return result 
    }
    
    q := &queue{}
    q2 := &queue{}
    
    q.enqueue(root)
    
    for q.isEmpty() != true {
        res := make([]int, 0)
        for q.isEmpty() != true {
            v := q.dequeue() 
            res = append(res, v.Val)
            q2.enqueue(v)
        }    
        
        for q2.isEmpty() != true {
            v := q2.dequeue()
            q.enqueue(v.Left)
            q.enqueue(v.Right)
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
    if node == nil {
        return
    }
    (*q) = append(*q, node)
}

func (q *queue) dequeue() *TreeNode {
    temp := (*q)[0]
    (*q) = (*q)[1:]
    fmt.Printf("returning: %v\n", temp)
    return temp 
}
