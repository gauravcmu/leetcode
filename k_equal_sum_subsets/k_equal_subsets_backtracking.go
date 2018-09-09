func canPartitionKSubsets(nums []int, k int) bool {
    sum := 0 

    for _, val := range nums {
        sum += val 
    }
    expected_sum := sum / k 
    fmt.Printf("1 sum:%v expected:%v\n",sum, expected_sum)
    if sum % k != 0 {
        return false
    }
    
    chosen := make([]int, k)
    return helper(nums, chosen, expected_sum, "  ")
}

func helper(nums []int, chosen []int, expected int, path string) bool {
    //fmt.Printf("%v-nums:%v chosen:%v expected:%v\n",path, nums, chosen, expected)
    if len(nums) == 0 {
        //base case. 
        for _,v := range chosen {
            if v != expected {
                return false 
            }
        }
        fmt.Printf("FOUND\n")
        return true
    } else {
        t := nums[len(nums)-1]
        nums = nums[:len(nums)-1]
        
        for index, v := range chosen {
            if v + t <= expected {
                chosen[index] += t
                found := helper(nums, chosen, expected, path+"  ")
                chosen[index] -= t
                if found == true {
                    return true
                }
            }    
        }    
        nums = append(nums, t)
    }
    return false
}
