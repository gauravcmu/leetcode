func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    i := 0
    j := 0
   
    m := len(nums1)
    n := len(nums2)

    var result []int 
    result = make([]int,0)
    
    for {
        if i == m || j == n {
            break
        }        
        if nums1[i] <= nums2[j] {
            result = append(result, nums1[i])
            i++
        } else {
            result = append(result, nums2[j])
            j++   
        }
    } 
    
    if i == m {
        for {
            result = append(result, nums2[j:] ...)
            break
        }
    }
    
    if j == n {
        for {
            result = append(result, nums1[i:] ...)
            break
        }
    }
    
    if (m+n)%2 == 0 {
        ix1 := float64(result[(m+n)/2])
        ix2 := float64(result[(m+n)/2 -1])
	    sum := (ix1 + ix2)
        return float64(sum/2.0)
    } else {
        return float64(result[(m+n)/2])
    }
}
