func wiggleSort(nums []int)  {
    for i:=0; i < len(nums);i++ {
        if i % 2 == 0 {
            if i+1 < len(nums) && nums[i] > nums[i+1] {
                   nums[i], nums[i+1] = nums[i+1], nums[i] 
                }
        } else {
            if i+1 < len(nums) {
                if nums[i] < nums[i+1] {
                   nums[i], nums[i+1] = nums[i+1], nums[i] 
                } 
            } 
        }        
    }
}