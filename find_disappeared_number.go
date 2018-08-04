func findDisappearedNumbers(nums []int) []int {
    res := make([]int, len(nums)+1)
    for _, v := range nums {
       res[v] = 1  
    }
    
    r := make([]int, 0)
    for i:=1; i<=len(nums);i++ {
        if res[i] == 0 {
            r = append(r, i)
        }
    }
    return r 
    
}
