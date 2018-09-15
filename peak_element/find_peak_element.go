func findPeakElement(nums []int) int {
    return getPeak(nums, 0, len(nums)-1) 
}

func getPeak(nums[]int, start, end int) int {
    mid := (start + end)/2  
    
    if (end - start == 0) {
        return start
    }
    
    if (end - start) == 1 {
        if nums[start] > nums[end] {
            return start 
        } else {
            return end
        } 
    }
    
    if (nums[mid-1] < nums[mid]) &&  (nums[mid] > nums[mid+1]) {
        return mid
    } 
    
    if nums[mid] < nums[mid-1] {
        return getPeak(nums, start, mid-1)
    } else if  nums[mid] < nums[mid+1] {
        return getPeak(nums, mid+1, end)
    }
    
    return -1 
}
