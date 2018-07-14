func reverseVowels(s string) string {
    bytes := []byte(s)
    
    j := len(bytes) - 1
    for i:= 0; i<j; {
        if isVowel(string(bytes[i])) {
           for j>i {
                if isVowel(string(bytes[j])) {
                    temp := bytes[i]
                    bytes[i] = bytes[j] 
                    bytes[j] = temp
                    j--
                    break
                } 
               j--
            } 
        } 
        i++
    }
    return string(bytes) 
}

func isVowel (s string) bool {
    var vowels map[string]bool = make(map[string]bool, 0 )  
    vowels["a"] = true 
    vowels["e"] = true
    vowels["i"] = true
    vowels["o"] = true
    vowels["u"] = true
    vowels["A"] = true 
    vowels["E"] = true
    vowels["I"] = true
    vowels["O"] = true
    vowels["U"] = true
    
    if _, ok:=  vowels[(s)]; ok {
        return true
    }  
    return false
}
