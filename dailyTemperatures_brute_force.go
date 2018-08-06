func dailyTemperatures(temperatures []int) []int {
    res := make([]int, 0)
    found := 0 
    for key, val := range temperatures {
        found = 0 
        for j:=key+1; j< len(temperatures); j++ {
            if temperatures[j] > val {
                res = append(res, j-key)
                found = 1
                break
            }
        } 
        if found == 0 {
            res = append(res, 0)
        }
    } 
    return res
}
