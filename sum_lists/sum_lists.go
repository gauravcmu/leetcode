/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var l3 *ListNode
    var units int = 0 
    var carry int = 0 
    for {
        if l1 == nil || l2 == nil  {
            break
        }

        units = (carry + l1.Val + l2.Val) % 10
        carry = (carry + l1.Val + l2.Val) / 10 
        addNode(&l3, units)
        l1 = l1.Next
        l2 = l2.Next 
    }
    
    if l1 == nil {
        for {
            if l2 == nil {
                break
            }
            
            units = (l2.Val + carry)%10          
            addNode(&l3, units)
            carry = (l2.Val+carry)/10
            l2 = l2.Next
        } 
    } else if l2 == nil {
         for {
            if l1 == nil {
                break
            }
            units = (l1.Val + carry)%10          
            addNode(&l3, units)
            carry = (l1.Val+carry)/10
            l1 = l1.Next
        }
    }
    if carry != 0 {
        addNode(&l3, carry)
    }
    return l3 
}

func addNode(l3 **ListNode, value int) {
    if *l3 == nil {
        *l3 = &ListNode {
            Val: value, 
            Next: nil, 
        }
        return 
    }
    
    temp := *l3 
    for {
        if temp.Next == nil {
            temp.Next = &ListNode {
            Val: value, 
            Next: nil ,
            }
            return
        } else {
                temp = temp.Next
            }
        }
        return 
}
