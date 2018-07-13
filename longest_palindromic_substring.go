func longestPalindrome(s string) string {
    longest := ""
    lenLongest := 0 
    start := 0 
    end := 0 
    if len(s) == 1 {
        return s
    }
    for i:=0; i< len(s); i++ {
        for j:=len(s)-1; j> i; j-- {
            if s[i] == s[j] {
                //potential palindrome in between i-j
                start = i 
                end = j 
                if isPalindrome(s, start, end) {
                    if end-start > lenLongest {
                        lenLongest = end-start
                        longest = s[start:end+1]
                    }
                }
            }
        } 
    }   
    if longest == "" {
        return string(s[0])
    }
    
    return longest 
}

func isPalindrome(s string, start, end int) bool {
    low := start
    high := end  
    
    for low < high {
        if s[low] == s[high] {
            low++ 
            high--
            continue
        } else {
            return false
        }
    }
    return true
}
