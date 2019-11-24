func nthUglyNumber(n int) int {
    count := 0 
    i := 1
    
    mapping := make(map[int]bool, 0)
    for i:=1; i<=6; i++ {
        mapping[i] = true
    }
    
    for i = 1; ; i++ {
        if isUgly(i, mapping) {
            count++
            if count == n {
                return i
            }
        } 
    }
    return -1
}

func isUgly(num int, mapping map[int]bool) bool {
    if num <= 0 {
        return false
    }

    return helper(num, mapping)
}


func helper(num int, mapping map[int]bool) bool {
    if num <= 6 {
        return true 
    }
 
    if _, ok := (mapping)[num]; ok {
        return (mapping)[num]
    } 
    
    if num % 2 == 0 {
        (mapping)[num] = helper(num/2, mapping)
        return (mapping)[num]
    }
    if num % 3 == 0 {
        (mapping)[num] = helper(num/3, mapping)
        return (mapping)[num]
    }
    if num %5 == 0 {
        (mapping)[num] = helper(num/5, mapping)
        return (mapping)[num]
    }
    
    (mapping)[num] = false
    return false 
}
