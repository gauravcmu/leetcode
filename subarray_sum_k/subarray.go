func subarraySum(nums []int, k int) int {
    mapping := make(map[int]int, 0)
    
    mapping[0] = 1 
    sum_till_index := 0 
    count := 0 
    
    for _, value := range nums {
        sum_till_index += value 
        if val, ok := mapping[sum_till_index - k]; ok {
            //found sum_till_index -k in map. thus a solution. Find how many such exist. 
            count = count + val // val has number of times the value (sum_till_index -k) is seen
        }
        if _, ok := mapping[sum_till_index]; ok {
            mapping[sum_till_index]++
        } else {
            mapping[sum_till_index] = 1 
        }
    }
    return count 
}
