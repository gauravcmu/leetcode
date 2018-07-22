func topKFrequent(nums []int, k int) []int {
    mapping := make(map[int]int, 0)
    
    for _, value := range nums {
        if _, ok := mapping[value]; ok {
            mapping[value] = mapping[value] + 1 
        } else {
            mapping[value] = 1  
        }
    }
    
    res := make([]kv, 0) 
    for key, value := range mapping {
        res = append(res, kv{key:key, val:value})
    }
    sort.Slice(res, func(a,b int)bool {
        return res[a].val < res[b].val 
    }) 
    
    fmt.Printf("res:%v\n", res)
    r := make([]int, 0)
    for i:=len(res)-1; i>len(res)-1-k;i-- {
        r = append(r, res[i].key) 
    }
    return r
}

type kv struct {
    key int 
    val int 
}
