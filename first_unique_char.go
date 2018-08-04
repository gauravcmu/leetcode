func firstUniqChar(s string) int {
    mapping := make(map[string]int, 0)
    for _, val := range s {
        if _, ok := mapping[string(val)]; ok {
            mapping[string(val)] += 1 
        } else {
            mapping[string(val)] = 1
        }
    }
    for index, v := range s {
        if mapping[string(v)] == 1 {
            return index 
        }
    }
    return -1 
}
