func longestConsecutive(nums []int) int {
    mapping := make(map[int]bool, 0)
    for _, value := range nums {
        mapping[value] = true 
    }
    
    // insert each element into hash table. 
    // iterate array and check each value in hash table. 
    // if the value-1 not found in hash table, then using this as start, find the consecutive elements in hash. 
    // maximize this.  
    var max int 
    for _, value := range nums {
        if _, ok := mapping[value-1]; !ok {
            count := 0 
           //first element. 
            for j := value; ; j++ {
                if _, ok := mapping[j] ; ok {
                    count++
                } else {
                    break
                }
            } 
            if max < count {
                max = count
            }
        }
    }
    return max 
}
