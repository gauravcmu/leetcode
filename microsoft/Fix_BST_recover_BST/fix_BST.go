package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
	var first *TreeNode
	var middle *TreeNode
	var last *TreeNode
	var prev *TreeNode

	findNodesToFix(root, &first, &middle, &last, &prev)

	fmt.Printf("first:%v middle:%v last:%v\n", first, middle, last)

	if first != nil && last != nil {
		temp := first.Val
		first.Val = last.Val
		last.Val = temp
	} else if middle != nil && first != nil {
		//second case, the two nodes appear next to each other in the in-order traversal
		temp := first.Val
		first.Val = middle.Val
		middle.Val = temp
	}
}


func findNodesToFix(root *TreeNode, first, middle, last, prev **TreeNode) {
	if root == nil {
		return
	}
	// fmt.Printf("first:%v middle:%v last:%v\n", *first, *middle, *last)

	// first find down the left subtree.
	findNodesToFix(root.Left, first, middle, last, prev)

	// if value at my node or value at root node is less than that of previous.
	//that means we found first discrepancy, so, previous is the first node that
	//is out of place. and this node is potentially the middle one (if two adjacent..)

	if *prev != nil && (*prev).Val > root.Val {
		//so we are looking for two out of place nodes. The two are lesser than previous ones.
		if *first == nil {
			*first = *prev
			*middle = root
		} else {
			*last = root
		}
	}

	//now since we are exploring right subtree, its previous will be the current root node.
	// fmt.Printf("first:%v middle:%v last:%v\n", *first, *middle, *last)

	*prev = root
	findNodesToFix(root.Right, first, middle, last, prev)
}