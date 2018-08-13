func productExceptSelf(nums []int) []int {
    l := len(nums)
    res := make([]int, l)
    
    productLeft := 1 
    for k, v := range nums {
        res[k] = productLeft 
        productLeft = productLeft * v  
    } 
    prodRight := 1
    for j := l-1; j>=0; j-- {
        res[j] = res[j] * prodRight 
        prodRight = prodRight * nums[j] 
    }
    
    return res  
}
