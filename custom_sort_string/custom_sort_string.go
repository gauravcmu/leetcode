import (
    "math"
    "sort"
)

//build s1 from S= {{val:a, index:2}, {val:b, index:}
//build s2 from T and use INT_MAX for index for values not in S. 
//use hash table for that 
//sort s2 based on index. 
//build a string from s2

type info struct {
	value rune 
    idx int 
} 
func customSortString(S string, T string) string {
	s1:= make([]info, 0) 
	mapping := make(map[string]info, 0)
	s2 := make([]info, 0)
	res := ""

	for key, val := range S {
        s1 = append(s1, info{value:val, idx:key})	
        mapping[string(val)] = info{value:val, idx:key}
    }
    for _, val := range T {
        if vmap, ok:= mapping[string(val)]; ok {
            s2 = append(s2, vmap)
        } else {
            s2 = append(s2, info{value:val, idx:math.MaxInt32})
        }
    }
    sort.Slice(s2, func(a, b int)bool {
        return s2[a].idx < s2[b].idx
    })

    for _,v := range s2 {
        res += string(v.value)
    }
    return res
}

