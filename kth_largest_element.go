var heapsize int 

func buildHeap(a[]int) {
    for i:=heapsize/2;i>=0;i-- {
        heapify(a, i)
    }
} 
func heapify(a[]int, index int) {
    left := 2*index
    right := 2*index +1
    var largest int

    if left <= heapsize && a[left] > a[index] {
       largest = left  
    } else {
        largest = index
    }
    if right <=heapsize && a[right] > a[largest] {
        largest = right
    }
    if largest != index {
        temp := a[index]
        a[index] = a[largest]
        a[largest] = temp
        
        heapify(a, largest)
    }
}
func extractMax(a[]int) int {
    if heapsize < 0 {
        return -1
    }
    max := a[0] 
    a[0] = a[heapsize]
    a[heapsize] = max
    heapsize = heapsize -1 
    heapify(a, 0)
    return max 
}

func findKthLargest(nums []int, k int) int {
    heapsize = len(nums) -1
    buildHeap(nums)  
    var temp int 
    for i:=0;i<k;i++ {
        temp = extractMax(nums)
    } 
    return temp
}
