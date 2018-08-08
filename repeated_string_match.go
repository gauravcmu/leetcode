/*
Given two strings A and B, find the minimum number of times A has to be repeated such that B is a substring of it. If no such solution, return -1.

For example, with A = "abcd" and B = "cdabcdab".

Return 3, because by repeating A three times (“abcdabcdabcd”), B is a substring of it; and B is not a substring of A repeated two times ("abcdabcd").

Note:
The length of A and B will be between 1 and 10000.


Solution:
    string a and string b are given. 
    n*a (a concatenated n times) -> contains string b
    n should be atleast cieling of lb/la 
    
    Find first occurence of first character of b in t (n*a)
    This is potential substring - check if it is substring. 
    if length of potential substring is less than b then add 1*A to t. 
    
    can optimize by moving to next start of A in t when not substring instead of going over it all
*/
func repeatedStringMatch(A string, B string) int {
    la := len(A)    
    lb := len(B)    
    if lb == 0 {
        return -1
    }
    n := (lb/la)+1
    t := "" 
    n1 := n
    for n1 > 0 {
        t += A 
        n1--
    }
    
    fmt.Printf("%v\n", t)
    for k, v := range t {
        if string(v) == string(B[0]) {
            //potential match starts here.
            if len(t[k:]) < lb {
                t+=A
                n++
            }
            if isSubstring(t[k:], B) {
                if k < (la) && k+lb > len(t)-la {
                    return n
                } else {
                    return n-1
                }
            }
        }
    }  
    return -1 
}

func isSubstring(a, b string) bool {
    if len(a) < len(b) {
        return false
    }
    for k, v := range b {
        if string(a[k]) == string(v) {
            continue
        }  else {
            return false
        }
    } 
    return true 
}
