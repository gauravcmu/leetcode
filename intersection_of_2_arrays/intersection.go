func intersection(nums1 []int, nums2 []int) []int {
    mapping := make(map[int]int, 0)
    res := make([]int, 0)
    
    //mark each element in the mapping 
    for _, v := range nums1 {
        if _, ok := mapping[v]; ok {
            mapping[v] += 1 
        } else {
            mapping[v] = 1 
        }
    }
    
    //saw a number in both, mark -1 
    for _, v := range nums2 {
        if _, ok := mapping[v]; ok {
            mapping[v] = -1 
        } 
    }
    //check for value as -1
    for key, val := range mapping {
        if val == -1 {
            res = append(res, key)
        }
    }
    return res 
}
