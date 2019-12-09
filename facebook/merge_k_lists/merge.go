import "container/heap"
type myheap []int

func (h myheap) Len() int {
	return len(h)
}

func (h myheap) Less(a, b int) bool {
	return h[a] < h[b]
}

func (h myheap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func (h *myheap) Push(val interface{}) {
	(*h) = append(*h, val.(int))
}

func (h *myheap) Pop() interface{} {
	t := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return t
}
func mergeKLists(lists []*ListNode) *ListNode {
	// create a heap.
	h := &myheap{}
	var last *ListNode
	var result *ListNode

	//iterate over lists. For each list, add the node to the priority queue.

	for {
		for index, list := range lists {
			if list != nil {
				heap.Push(h, list.Val)
				list = list.Next
                lists[index] = list 
			}
		}

		if h.Len() == 0 {
			break
		} else {
			//pop the root element and push it down the linked list.
			v := heap.Pop(h)
			if result == nil {
                result = insertToList(last, v.(int))
				last = result
			} else {
                last = insertToList(last, v.(int))
			}
		}
	}
    return result 
}

//insert node to the end of the list, where last is passed in.
func insertToList(last *ListNode, value int) *ListNode {
	if last == nil {
		return &ListNode{
			Val:  value,
			Next: nil,
		}
	} else {
		last.Next = &ListNode{
			Val:  value,
			Next: nil,
		}
		return last.Next
	}
}
