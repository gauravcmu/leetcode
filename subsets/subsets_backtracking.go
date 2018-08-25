func subsets(nums []int) [][]int {
    result := make([][]int, 0)
    subsetsHelper(nums, []int{}, &result)
    return result
}

func subsetsHelper(nums []int, chosen []int, result *[][]int) {
    if len(nums) == 0 {
        (*result) = append(*result, chosen)
    } else {
        //remove last element from nums. 
        //explore 
        //add last removed element to chosen and explore.
        temp := nums[:len(nums)-1]

        subsetsHelper(temp, chosen, result)
        
        chosen = append(chosen, nums[len(nums)-1])
        subsetsHelper(temp, chosen, result)
        
        chosen = chosen[:len(chosen)-1] 
    } 
}
