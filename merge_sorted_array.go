func merge(nums1 []int, m int, nums2 []int, n int)  {
    //do merge and use end of array as 3rd array. Merge biggest element from m+n -1 index..
    k := m+n-1
    i := m-1
    j := n-1
    
//[1,2,0,3,5,6]
//[2,0,0]
    
    for ; i>=0 && j>=0; {
        if nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
            k--
        }  else if nums1[i] < nums2[j] {
            nums1[k] = nums2[j]
            j--
            k--
        } else {
            //equal
            fmt.Printf("nums1:%v nums2:%v\n", nums1,nums2)
            
            nums1[k] = nums1[i] 
            i--
            k--
            fmt.Printf("nums1:%v nums2:%v\n", nums1,nums2)
            
            nums1[k] = nums2[j] 
            j--
            k--
            fmt.Printf("nums1:%v nums2:%v\n", nums1,nums2)
        }
    }
    
    for i >= 0 && k >= 0 {
        fmt.Printf("2 nums1:%v nums2:%v i:%v k%v \n", nums1,nums2, i, k)
        nums1[k] = nums1[i]
        i--
        k--
    }
    for j >= 0 && k >= 0 {
        fmt.Printf("3 nums1:%v nums2:%v\n", nums1,nums2)
        nums1[k] = nums2[j]
        j--
        k--
    }
}

