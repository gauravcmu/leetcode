/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var res *ListNode = nil 
    
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            t := l1.Next 
            res = insertNode(res, l1)
            l1 = t 
        } else if l2.Val < l1.Val {
            t := l2.Next 
            res = insertNode(res, l2)
            l2 = t 
        } else {
            t := l1.Next 
            res = insertNode(res, l1)
            l1 = t
            
            t = l2.Next 
            res = insertNode(res, l2)
            l2 = t  
        }
    } 
    
    for l1 != nil {
            t := l1.Next 
            res = insertNode(res, l1)
            l1 = t 
    }
    
    for l2 != nil {
            t := l2.Next 
            res = insertNode(res, l2)
            l2 = t 
    }
    return res  
}
//todo - keep track of last inserted node for faster inserts
func insertNode(head *ListNode, node *ListNode) *ListNode {
    temp := head
    if head == nil {
        head = node
        head.Next = nil 
        return head
    }
    
    for head.Next != nil {
        head = head.Next 
    }
    
    head.Next = node
    head.Next.Next = nil 
    return temp   
}

