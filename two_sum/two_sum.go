type data struct {
    value int 
    index int 
}
func twoSum(nums []int, target int) []int {
    hash := make(map[int][]data, 0)  
    
    for key, value := range nums {
        if _, ok := hash[value]; !ok {
            hash[value] = make([]data,0)
        }
        hash[value] = append(hash[value], data{value: value, index: key, }) 
        if k, ok := hash[(target - value)]; ok {
            for _, data := range k {
                if data.index == key {
                    continue
                } else {
                    return []int{key, data.index}
                }
            }
        } 
    }
    return []int{-1,-1}
}
