func subsets(nums []int) [][]int {
    res := make([][]int, 0)
    chosen := make([]int, 0) 
    sh(&res, nums, chosen)  
    return res 
}

func sh(res *[][]int, nums[]int, chosen []int) {
    if len(nums) == 0 {
        *res = append(*res, chosen) 
    } else {
        t := nums[len(nums)-1] 
        nums = nums[:len(nums)-1]
        
        //once a member is removed from nums and not chosen either..explore. 
        sh(res, nums, chosen) 
         
        //member removed from nums and chosen, now explore
        chosen = append(chosen, t)
        sh(res, nums, chosen) 
        
        //time to unchose
        nums = append(nums, t)
        chosen = chosen[:len(chosen)-1]
    }
}
