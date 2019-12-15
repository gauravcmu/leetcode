package main

func main() {
	list := flatten_multilevel_list(list )
}

type node struct {
	int value
	next *node
	child * node
}

//	1		->2		->3
//child	 	21		-> 9
//child	 	212
/*
	flatten(1)
		temp is at 2.
			flatten(21)
				temp is at 21 - next: 9, -> 21 -> 212 -> 9
				flatten(212) - return 212, 212
					- prev points to 212.

 */

func flatten_multilevel_list(list * node) *node {
	l, _ := flatten(list)
	return l
}

func flatten(list *node) (*node, *node) { //return head and last node?
	if list == nil {
		return nil, nil
	}

	temp := list
	var prev *node
	for ; temp != nil; {
		prev = temp
		if temp.child != nil {
			//flatten child list.
			//that should get us the child list head and tail.
			// for this node, save my next and add child linked list here.
			next := temp.next
			child_head, child_last := flatten(temp.child)
			temp.next = child_head
			child_last = next
		}
		temp = temp.Next
	}
	//prev is last node as temp went to nil and exited the loop.
	return list, prev
}