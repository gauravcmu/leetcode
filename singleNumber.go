func singleNumber(nums []int) int {
    var res int 
    for _, val := range nums {
        res ^= val  
    }   
    return res
}
Given a non-empty array of integers, every element appears twice except for one. Find that single one.
