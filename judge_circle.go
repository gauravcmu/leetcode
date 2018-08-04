func judgeCircle(moves string) bool {
    mapping := make(map[string]int, 0)   
    
    for _, value := range moves {
        if _, ok := mapping[string(value)]; ok {
            mapping[string(value)] +=1
        } else {
            mapping[string(value)] = 1
        }
    }
    
    if mapping["U"] != mapping["D"] {
        return false
    }
    if mapping["L"] != mapping["R"] {
        return false
    }
    return true
}
