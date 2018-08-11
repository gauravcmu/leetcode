func countBits(num int) []int {
    res := make([]int, 0)
    for i:=0;i<=num; i++ {
        res = append(res, numBits(i))
    }  
    return res
}
func numBits(n int) int {
    if n == 0 {
        return 0
    }
    count := 0 
    for (n > 0) {
        n = n&(n-1)
        count++
    }
    return count
}
