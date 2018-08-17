// spaces part of palindrome? - ignored. 
// Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
import "strings"

//builds alphan and returns lower case
func buildAlphaN(s string) string {
    res := ""
    for _, v := range s {
        if (v >= 48 && v <=57) || (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
            res += string(v)
        } 
    } 
        
    return strings.ToLower(res) 
}

func isPalindrome(s string) bool {
    if len(s) == 0 {
        //empty string
        return true
    } 
    
    s2 := buildAlphaN(s)
    fmt.Printf("s:%v s2:%v\n", s, s2)
    
    if palindromeCheck(s2) == true {
        return true
    }
    return false 
}

func  palindromeCheck(s string) bool {
    if len(s) == 1 {
        return true
    }
    
    for i, j := 0, len(s) - 1; i < j; i,j = i+1, j-1 {
        if s[i] == s[j] {
            continue
        } else {
            return false 
        } 
    }  
    return true
}
