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
    //biggest perfect square
    bps := int(math.Sqrt(float64(n)))
    val := bps
    minSum := n 
    sum := 0
    if val != 0 {
        for (val >= 1) {
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
