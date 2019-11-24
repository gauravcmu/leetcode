func isUgly(num int) bool {
    if num <= 0 {
        return false
    }
    mapping := make(map[int]bool, 0)
    return helper(num, &mapping)
}


func helper(num int, mapping *map[int]bool) bool {
    if num <= 6 {
        return true 
    }
     
    if _, ok := (*mapping)[num]; ok {
        return (*mapping)[num]
    }
    
    if num % 2 == 0 {
        (*mapping)[num/2] = helper(num/2, mapping)
        return (*mapping)[num/2]
    }
    if num % 3 == 0 {
        (*mapping)[num/3] = helper(num/3, mapping)
        return (*mapping)[num/3]
    }
    if num %5 == 0 {
        (*mapping)[num/5] = helper(num/5, mapping)
        return (*mapping)[num/5]
    }
    
    return false 
}
