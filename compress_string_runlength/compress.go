import "strconv"
func compress(chars []byte) int {
    charstemp := compress_helper(chars)
    return len(charstemp)
}

//aaabbbccc = "a3" + compress_helper("bbbccc")
func compress_helper(chars []byte) [] byte {
    if len(chars) <= 1 {
        //base case
        return chars 
    }
    
    w := chars[0]
    var count int = 1  
    var i int 
    //get repetitions for the character
    for i=1; i<len(chars);i++ {
        if chars[i] == w {
            count++
        } else {
            break
        } 
    }    
    var res []byte = make([]byte, 0) 
    if count == 1 {
        //a
        res = append(res, w)
    } else {
        //a3
        res = append([]byte{w}, []byte(strconv.Itoa(count))...)
    }
    //make recursive call to compress rest of the string
    res = append(res, compress_helper(chars[i:])...)
    for k, v := range res {
        chars[k] = v
    }
    
    chars = chars[:len(res)]   
    return chars     
}
