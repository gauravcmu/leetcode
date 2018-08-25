func numTrees(n int) int {
    var cache map[int]int = make(map[int]int, 0)
    return catalan(n, &cache)
}

func catalan(n int, cache *map[int]int) int {
    var sum int = 0 
    
    //cat(n) = for i (0 to n-1), sum of cat(i) * cat(n-1-i) 
    if n <= 1 {
        return 1
    }
    
    if v, ok := (*cache)[n]; ok {
        //cache lookup
        return v
    } 
    
    for i:=0; i<n; i++ {
        sum += catalan(i, cache)*catalan(n-1-i, cache)
    }
    
    (*cache)[n] = sum
    fmt.Printf("setup cache:%v\n", *cache)
    return sum
}
