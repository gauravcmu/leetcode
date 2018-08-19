/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


/*
 * Push first half of the list on a stack. 
 * ignore center element if odd number list 
 * pop and compare second half with stack      
 */
func isPalindrome(head *ListNode) bool {
    stack := make([]int, 0)
    if head == nil {
        return true
    } 
    if head.Next == nil {
        return true
    } 
    
    l := findLength(head) 
    
    temp := head
    for i:=0; i<l/2; i++ {
        push(&stack, temp.Val)
        temp = temp.Next 
    }
    if l%2 != 0 {
        //odd number of nodes. Skip one (middle element)
        temp = temp.Next
    }
    
    fmt.Printf("stack:%v temp:%v\n",stack, temp)
    for (temp != nil)  {
        if (isEmpty(&stack)) {
            return false 
        }
        if temp.Val == pop(&stack) {
            temp = temp.Next
            continue
        } else {
            //found non matchin element, false. 
            return false
        }
    }
    return isEmpty(&stack) 
}

func findLength(head *ListNode) int {
    var count int
    for (head != nil) {
        head = head.Next
        count++ 
    }
    return count 
}

func isEmpty(stack *[]int) bool {
    return len(*stack) == 0
}
func push(stack *[]int, val int) {
    *stack = append(*stack, val)
}

func pop(stack *[]int) int {
    t := (*stack)[len(*stack)-1] 
    *stack = (*stack)[:len(*stack)-1] 
    return t 
}

func peek(stack *[]int) int {
    return (*stack)[len(*stack)-1] 
}
