func anagramMappings(A []int, B []int) []int {
    res := make([]int, 0)
    for _,v := range A {
        for index, v2 := range B {
            if v == v2 {
                res = append(res, index)  
                break
            }
        }
    }
    return res
}
