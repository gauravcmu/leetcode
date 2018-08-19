func canPermutePalindrome(s string) bool {
    //permutation of palindrome if 
    //is anagram of a palindrome
    //for palindrome - 
    // if length % 2 == 0 then even number of each character. 
    // else even number of each character but one. 

    mapping := make(map[string]int, 0)
    for _, v := range s {
        if _, ok := mapping[string(v)]; !ok {
            mapping[string(v)] = 1 
        } else{
            mapping[string(v)] += 1 
        }
    }
    
    var foundOdd bool 
    for _,v := range mapping {
        if v % 2 != 0 {
            if foundOdd == false {
                foundOdd = true
            } else {
                //there should be exactly one char with exactly 1 occurrence
                return false
            }
        } 
    }
    return len(s) % 2 == 0 && foundOdd == false || len(s) % 2 != 0 && foundOdd == true 
}
