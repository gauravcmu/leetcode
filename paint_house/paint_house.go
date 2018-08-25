import "math"

func minCost(costs [][]int) int {
    min := math.MaxInt32 
    minCostHelper(costs, []int{}, &min,-1, "  ")
    return min 
}

func minCostHelper(available[][]int, chosen[]int, min *int, lastindex int, indent string) {
    var sum int = 0 
    //fmt.Printf("%vavail: %v chosen:%v min:%v lastindex:%v\n", indent, available, chosen, min, lastindex)
    
    if len(available) == 0 {
        for _, v := range chosen {
            sum += v          
        } 
        if sum < *min {
            *min = sum 
        }
     //   fmt.Printf("Min:%v\n", *min)
    } else {
        for i:=0; i<len(available[0]); i++ {
            //[[17,2,17],[16,16,5],[14,3,19]]
                //i = 0 -> [[16,16,5],[14,3,19]] [17] lastindex = 0 
            if i == lastindex {
                continue 
            }
            if available[0][i] > *min {
               continue 
            }
            //chose 
            chosen = append(chosen, available[0][i])
            //explore
            if len(available) >= 1 {
                minCostHelper(available[1:], chosen, min, i, indent+"  ")
            }
            //unchose
            chosen = chosen[:len(chosen)-1]
        }          
    
    } 
    
    
}
