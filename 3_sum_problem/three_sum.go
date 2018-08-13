func threeSum(nums []int) [][]int {
    result := make([][]int, 0)
    sort.Slice(nums, func (i, j int) bool{
        return nums[i] < nums[j]
    }) 
   // [ -2 -2 -1 0 1 2 3 4 5 ] 
   //    ^  ^ --->    <--- ^ 
   //    i  s              e 
    for i:= 0; i< len(nums)-2; i++ {
        if (i==0 || nums[i] > nums[i-1]) {
            start := i+1
            end := len(nums)-1
            
            for start < end {
                if (nums[i]+nums[start]+nums[end] == 0) {
                    r := make([]int, 0)
                    r = append(r, nums[i], nums[start], nums[end])
                    result = append(result, r)
                }
                
                if (nums[i] + nums[start] + nums[end]) < 0 {
                    currentstart := start
                    for ((nums[currentstart] == nums[start]) && (start < end)) {
                           start++ 
                    }
                } else {
                    currentend := end 
                    for ((nums[currentend] == nums[end]) && (start < end)) {
                           end-- 
                    }
                }
            }
        }
    }

   return result
}
