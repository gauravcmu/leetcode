/*
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6

Solution:
	1. Go over each list one by one, one node at a time. 
	2. Insert this node to a min Heap
	3. At the end of 1. and 2. Heap contains all front nodes from all lists. 
	4. get minimum element from heap and insert to linked list (save end node pointer for quick insertion to ll)

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "container/heap"

type IntHeap []int 

func (h IntHeap) Len() int {
    return len(h)
}

func (h IntHeap) Less(a, b int) bool {
    return h[a]  < h[b]
}

func (h IntHeap) Swap(a, b int) {
    h[a], h[b] = h[b], h[a]
}

//container.heap.Push() calls this method and heapifies the heap after
//Call heap.Push() and never IntHeap.Push().
func (h *IntHeap) Push (value interface{}) {
    *h = append(*h, value.(int))
}
//container.heap.Pop() does magic - swaps 0th element with nth. 
//Pop thus should remove nth element and adjust the slice. 
func (h *IntHeap) Pop() interface{} {
    l := len(*h) 
    res := (*h)[l-1] 
    *h = (*h)[0:l-1]
    return res
}

func mergeKLists(lists []*ListNode) *ListNode {
    h := &IntHeap{} 
    heap.Init(h) 
    temp := make([]*ListNode,len(lists))
    var result *ListNode 
    var last *ListNode 
    
    for _, l := range lists {
        temp = append(temp, l) 
    }
    
    for { 
        //iterate over lists array 
        for index, l := range temp {
            if l != nil {
                heap.Push(h, l.Val)
                l = l.Next
                temp[index] = l 
            }
        } 
        if h.Len() == 0 {
            //all done. Nothing in priority queue 
            break
        } else {
            node := heap.Pop(h)
            if result == nil {
                //preserve head of list
                result = insertToList(result, node.(int))
                last = result 
            } else {
                last = insertToList(last, node.(int)) 
            }
        }
    }
    return result 
}

//insert and return 
func insertToList(node *ListNode, Val int) (lastInList *ListNode) {
    if node == nil {
        //create a node
        n := &ListNode{
            Val: Val,
            Next: nil, 
        }
        fmt.Printf("Added 1:%+v\n", n)
        return n 
    }  else {
        node.Next = &ListNode{
            Val: Val,
            Next: nil, 
        }
        fmt.Printf("Added 2:%+v\n", node.Next)
        return node.Next
    }
    return nil 
}
