package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// go over l1 and l2 and break if either one is nil
	// go over l1 if non nil and add carry and nodes to result
	// same as above for l2.
	carry := 0

	var head *ListNode
	for l1 != nil && l2!=nil {
		digit := (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry ) / 10
		addNode(&head, digit)
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		digit := (l1.Val + carry) % 10
		carry = (l1.Val + carry ) / 10
		addNode(&head, digit)
		l1 = l1.Next
	}
	for l2 != nil {
		digit := (l2.Val + carry) % 10
		carry = (l2.Val + carry ) / 10
		addNode(&head, digit)
		l2 = l2.Next
	}

	if carry != 0 {
		addNode(&head, carry)
	}
	return head
}


func addNode(head **ListNode, val int) *ListNode {
	fmt.Printf("addnode(): %v: %v", head, val)
	if *head == nil {

		*head = &ListNode{
			Val: val,
			Next: nil,
		}
		return *head
	}

	t := *head
	for t != nil && t.Next!=nil {
		t = t.Next
	}

	t.Next = &ListNode{
		Val: val,
		Next: nil,
	}
	return *head
}
