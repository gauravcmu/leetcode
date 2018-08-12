/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

func merge(intervals []Interval) []Interval {
    if len(intervals) <= 1 {
        return intervals
    }
    
    res := make([]Interval, 0)
    sort.Slice(intervals, func(a,b int) bool {
        return intervals[a].Start < intervals[b].Start
    })

    //add first element to result
    res = append(res, intervals[0])
    for key, value := range intervals {
        if key == 0 {
            continue
        }
        mergeWithRes(&res, value)
     }	
    return res
}

func mergeWithRes(res *[]Interval, value Interval) {
	for key, v := range *res {
		f, s, e := isOverLapping(v.Start, v.End, value.Start, value.End)
            if f == true {
                fmt.Printf("value:%v res:%v\n", value, res)
                v.Start = s
                v.End = e 
                (*res)[key] = v
                return 
            } 
    }
    *res = append(*res, value)
}				
             //      1        6     8         10 
func isOverLapping(aStart, aEnd, bStart, bEnd int) (bool, int, int) {
    if aStart > bEnd || bStart > aEnd {
        return false, 0, 0 
    }
    if aStart <= bStart { 
        if aEnd >= bEnd { 
            return true, aStart, aEnd
        } else {		
            return true, aStart, bEnd
        }
    }
    if bStart < aStart {
        if bEnd >= aEnd {
            return true, bStart, bEnd
        } else {
            return true, bStart, aEnd
        }
    }
    return false, 0,0
}

