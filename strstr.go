func strStr(haystack string, needle string) int {
    l1 := len(haystack)
    l2 := len(needle)
    if l2 > l1 {
        return -1
    }
    if l1 == 0 || l2 == 0 {
        return 0
    }
    
    for i:=0; i<l1; {
            if string(haystack[i]) == string(needle[0]) {
                if i+l2 <= l1 && haystack[i+l2-1] == needle[l2-1] {
                    //try doing match
                    match, next := isMatch(needle, haystack[i:i+l2]) 
                    if match == true {
                        return i 
                    } else {
                        if next != -1 {
                            i += next 
                            continue
                        } 
                    }
                } 
            } 
        i++
    } 
    return -1 
}

//find match else suggest next starting point
func isMatch(a, b string) (ismatch bool, nxt int) {
    var match bool = true  
    var next int = -1 
    for k,_ := range a {
        if a[k] != b[k] {
            match = false 
        }   
        //save a potential next starting point
        //ensure we don't overwrite it (next should be -1)
        if next == -1 && b[k] == a[0] && k!=0 {
            next = k
        }
    }
    
    if match == false {
        return false, next
    }
    
    return true, 0
}
