func nthUglyNumber(n int) int {
    index2 := 0 
    index3 := 0 
    index5 := 0 
    
    result := make([]int, 0)
    result = append(result, 1)
    
    
    for i := 1; i <= n;i++ {
        prod2 := result[index2] * 2 
        prod3 := result[index3] * 3 
        prod5 := result[index5] * 5 
        
        v := min(prod2, prod3, prod5)
        if v == prod2 {
            index2++
        }  
        if v == prod3 {
            index3++ 
        } 
        if v == prod5 {
            index5++
        }
        result = append(result, v)        
    }
    return result[n-1]
}

func min(a, b, c int) int {
    if a < b {
        if a < c {
            return a 
        } else {
            return c 
        }
    } else {
          if b < c {
            return b
        } else {
            return c 
        }
    }
}
