import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	temp := head
	count := 0

	mapping := make(map[int]*ListNode, 0)

	for temp != nil {
		count++
		mapping[count] = temp
		temp = temp.Next
	}
	fmt.Printf("%+v\n\n", mapping)

	return sortedListToBSTHelper("  ", head, 1, count, mapping)
}

func sortedListToBSTHelper(path string, head *ListNode, start, end int, mapping map[int]*ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	fmt.Printf("%v head:%+v start:%v end:%v \n", path, head, start, end)
	if start == end {
		newNode := &TreeNode{Val: head.Val, Left: nil, Right: nil}
		return newNode
	}

	if start < end {
		mid := start + (end-start)/2

		midNode := mapping[mid]
		newNode := &TreeNode{Val: midNode.Val}

		newNode.Left = sortedListToBSTHelper(path+"  ", head, start, mid-1, mapping)
		newNode.Right = sortedListToBSTHelper(path+"  ", midNode.Next, mid+1, end, mapping)
		return newNode
	}
	return nil
}