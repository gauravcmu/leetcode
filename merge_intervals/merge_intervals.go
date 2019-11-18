func merge(intervals [][]int) [][]int {
    result := make([][]int, 0)
    if len(intervals) == 0 {
        return result 
    }
    
    sort.Slice(intervals, func(a, b int) bool {
        return intervals[a][0] < intervals[b][0]
    })
    
    result = append(result, intervals[0])
    
    for i, j := 0, 1; j < len(intervals); j++ {
        if intervals[j][0] <= result[i][1] {
            if intervals[j][1] <= result[i][1] {
                continue
            } else {
                result[i][1] = intervals[j][1]
            }
        } else {
            i++
            result = append(result, intervals[j])
        }
    }
    return result 
}
