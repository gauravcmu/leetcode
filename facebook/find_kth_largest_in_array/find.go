import "container/heap"

func findKthLargest(nums []int, k int) int {
    h := &myheap{}
    var v int 
    for _, value := range nums {
        heap.Push(h, value) 
    }
    
    for i := 0; i < k; i++ {
        v = heap.Pop(h).(int)
    }
    
    return v 
}


type myheap []int 

func (h myheap)Len() int {
    return len(h)
}

func (h myheap) Less(a, b int) bool {
    return h[a] > h[b]
}

func (h myheap) Swap(a, b int) {
    h[a], h[b] = h[b], h[a]
}

func (h * myheap) Push(val interface{}) {
    (*h) = append(*h, val.(int))
}

func (h* myheap) Pop() interface{} {
    t := (*h)[len(*h)-1]
    (*h) = (*h)[:len(*h)-1]
    return t 
}

