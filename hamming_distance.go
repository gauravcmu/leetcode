func hammingDistance(x int, y int) int {
    s := x ^ y
    count := 0
    
    for s != 0 {
        s = s&(s-1)
        count++
    } 
    return count
}
