func plusOne(digits []int) []int {
    res := make([]int, len(digits)+1)
    l := len(digits)

    if digits[l-1] < 9 {
        digits[l-1]+=1
        return digits 
    }  
    
    var val int 
    var carry int 
    
    val =  (carry + digits[l-1]+ 1) % 10  
    carry = (carry + digits[l-1]+ 1) / 10
    res[l] = val 
    
    for i:=len(digits) -2; i>=0; i-- {
        val =  (carry + digits[i]) % 10  
        carry = (carry + digits[i]) / 10
        res[i+1] = val 
        fmt.Printf("val:%v carry:%v res:%v\n", val, carry, res)
    }
    
    if carry != 0 {
        res[0] = carry   
        return res 
    } 
    return res[1:]
}
