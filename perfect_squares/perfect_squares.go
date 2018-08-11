import "math"
func min (a, b int) int {
    if a < b {
        return a
    }
    return b
}

func numSquares(n int) int {
    cache := make(map[int]int, 0)
    return cachedNumSquares(n, cache)
}

// func(n) = 1 if perfect square or func(perfect square) + func (n-perfectS*perfectS)
func cachedNumSquares(n int, cache map[int]int) int {
    if n <= 1 {
        return n  
    }
    if v, ok := cache[n]; ok {
        return v
    } 
    //biggest perfect square root
    bps := int(math.Sqrt(float64(n)))
    val := bps
    minSum := n  // 1 added n times. 
    sum := 0
    if val != 0 {
        for (val >= 1) {
	    //go from bggest perfect square to 1 and minimize sum 
	    //recursively do func(perfect square) + func (n-perfectS*perfectS)
            sum = cachedNumSquares(n - val*val, cache) + 1 
            if sum < minSum {
                minSum = sum
            } 
            val--
        } 
    }  
    cache[n] = minSum
    return minSum  
}
