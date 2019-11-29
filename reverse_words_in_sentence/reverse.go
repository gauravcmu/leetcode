import "strings"

/* 
input: "the sky   is blue"
               ^  ^
get input as byte array. 
in = []byte(s)

1.  First pass - get tokens of the string. 
    reverse the token array. 
2.  add spaces and return string. 
*/

func reverseWords(s string) string {
    //Get tokens from input string 
    tokens := strings.Fields(s)
    res := make([]byte, 0)
    
    fmt.Printf("tokens: %v.\n", tokens)
    
    //reverse token array
    for i, j :=0, len(tokens)-1; i<j;i, j = i+1, j-1 {
        tokens[i], tokens[j] =  tokens[j], tokens[i]
    }
    
    fmt.Printf("tokens: %v.\n", tokens)
    
    //Add space after each field of token array.
    for index, token := range tokens {
        res = append(res, []byte(token)...)
        if index != len(tokens)-1 {
            res =append(res, ' ')
        }
    }
    return string(res)
}

