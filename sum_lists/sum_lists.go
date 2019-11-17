/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var result *ListNode = nil 
    carry := 0 
    value := 0 
    
    for {
        if l1 == nil || l2 == nil {
            break 
        } else {
            sum :=  (carry + l1.Val + l2.Val)
            if sum >= 10 {
                value = (sum)%10
                carry = (sum)/10   
            } else {
                value = sum  
            }
        }
        result = addNode(result, value)
        l1 = l1.Next
        l2 = l2.Next
    }

    if l1 == nil {
        for l2 != nil {
            sum := carry + l2.Val 
            if sum >= 10 {
                value = (sum)%10
                carry = (sum)/10
            } else {
                value = sum
            }

            result = addNode(result, value)
            l2 = l2.Next
        }
    }
    
    if l2 == nil {
        for l1 !=nil {
            sum := carry + l1.Val 
            if sum >= 10 {
                value = (sum)%10
                carry = (sum)/10
            } else {
                value = sum
            } 
           result = addNode(result, value)
        }

    }
    
    if carry != 0 {
        result = addNode(result, carry)
    }
    
    return result     
}

func addNode(list *ListNode, value int) *ListNode {     
    fmt.Printf("adding value:%v to list:%+v\n", value, list)
    node := ListNode {
        Val: value, 
        Next: nil, 
    } 
    
    if list == nil {
        return &node
    }
    
    temp := list 
    for ; temp.Next != nil; temp = temp.Next {
    }
    temp.Next = &node 
    
    fmt.Printf("Added: %v to list: %+v\n", value, list)  
    return list 
}
