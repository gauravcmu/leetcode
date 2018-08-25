import "sort" 

func groupAnagrams(strs []string) [][]string {
    mapping := make(map[string][]string,0)
    
    for _, v := range strs {
        //go over string by string. 
        runes := []rune(v)
        sort.Slice(runes, func(a,b int) bool {
            return runes[a] < runes[b]
        })
        word := string(runes) 
        
        if _, ok := mapping[word]; !ok {
            mapping[word] = make([]string, 0)
        } 
        mapping[word] = append(mapping[word], v)
    }
    
    result := make([][]string, 0)
    for _, list := range mapping {
        result = append(result, list)
    }
    return result
}
