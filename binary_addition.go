func addBinary(a string, b string) string {
    s1 := a
    s2 := b
    res := ""
    
    carry := "0" 
    i := len(a) -1 
    j := len(b) -1
    var v string 
    if a == "0" {
        return b
    } else if b == "0" {
        return a
    }
    
    for ; (i>= 0)&&(j >= 0); i,j=i-1,j-1 {
        v, carry = addDigits(string(s1[i]), string(s2[j]), carry) 
        res = v + res 
    }   
    
    for ; i >= 0;i = i-1 {
        v, carry = addDigits(string(s1[i]), "0", carry) 
        res = v + res 
    }   
    
    for ; j >= 0;j = j-1 {
        v, carry = addDigits("0", string(s2[j]), carry) 
        res = v + res 
    }   
    
    if carry != "0" {
         res = carry + res
    }
    return res
        
}
func addDigits(a, b, carry string) (val, car string) {
    if carry == "0" {
        if (a == "1") && (b == "1") {
            return "0", "1"
        }
        if a == "1" && b == "0" {
            return "1", "0"
        }
        if a == "0" && b == "0" {
            return "0", "0"
        }
        if a == "0" && b == "1" {
            return "1", "0"
        }
    } else {
        if a == "1" && b == "1" {
            return "1", "1"  
        }
        if a == "1" && b == "0" {
            return "0", "1" 
        }
        if a == "0" && b == "0" {
            return "1", "0"
        }
        if a == "0" && b == "1" {
            return "0", "1" 
        }
    }
    return "0","0"
}
