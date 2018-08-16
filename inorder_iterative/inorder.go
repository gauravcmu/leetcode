/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    result := make([]int,0)
    stack := make([]*TreeNode,0)
    if root == nil {
        return []int{}
    }
     
    visited := make(map[*TreeNode]bool, 0)
    
    push(&stack, root) 

    for !isEmpty(stack) {
        v := peek(stack) 
        
        if (v.Left != nil) && (visited[v.Left] != true) {
            //left node not visited..push
            push(&stack, v.Left) 
        } else {
            //handle the node / add to result
            result = append(result, v.Val)
            pop(&stack)
            visited[v] = true
            if (v.Right != nil) && (visited[v.Right] != true) {
                push(&stack, v.Right)
            } 
        }
        
    }
    return result 
    
}
                            
func push(x *[]*TreeNode, val *TreeNode) {
	*x = append(*x, val)
}
func pop(x *[]*TreeNode) *TreeNode {
	val := (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return val
}
func peek (x []*TreeNode) *TreeNode {
	return x[len(x)-1]
}

func isEmpty(x []*TreeNode) bool {
	return len(x) <= 0
}

