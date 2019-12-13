package loop_ll
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


//start with fast pointing to next node of head.
//slow pointing to head.
//move fast two places and slow one.
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	slow := head


	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
