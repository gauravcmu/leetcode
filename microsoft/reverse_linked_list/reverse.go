package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	return revList(head)
}

func revList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	rest := revList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rest
}

//iterative
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	temp := head.Next
	var prev *ListNode

	for temp != nil && temp.Next != nil {
		t := temp.Next
		temp.Next = head
		head.Next = prev
		prev = head
		head = temp
		temp = t
	}

	temp.Next = head

	return temp
}
