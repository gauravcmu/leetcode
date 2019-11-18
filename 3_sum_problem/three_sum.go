func threeSum(nums []int) [][]int {
    result := make([][]int, 0 )
    
    sort.Slice(nums, func(a, b int)bool {
        return nums[a] < nums[b]
    })
    
    for i:=0; i < len(nums)-2; i++ {
        if i==0 || nums[i] > nums[i-1] { 
            start := i+1
            end := len(nums)-1
            
            for start < end {
                sum := nums[i] + nums[start]+ nums[end]
                if sum == 0 {
                    result = append(result, []int{nums[i], nums[start], nums[end]})
                }
                if sum > 0 {
                    currentend := end 
                    for nums[currentend] == nums[end] && start < end {
                        end--
                    }
                } else {
                    currentstart := start
                    for nums[currentstart]==nums[start] && start < end {
                        start++ 
                    }
                }
            }
        }
        
    }
    return result 
}
