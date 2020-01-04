package pour_water

import "fmt"

func pourWater(heights []int, V int, K int) []int {
	if len(heights) == 0 || K >= len(heights){
		return heights
	}

	for count := 0; count < V; count++ {
		if K > 0 && heights[K-1] < heights[K] {
			old := K
			for j := K - 1; j >= 0 && heights[j] < heights[old]; j-- {
				old = j
			}
			//old should be the last hieght seen where we can add the drop
			fmt.Printf("adding drop at:%v\n", old)
			heights[old]++
		} else if K < len(heights)-1 && heights[K] > heights[K+1] {
			old := K
			for j := K + 1; j < len(heights) && heights[j] < heights[old]; j++ {
				old = j
			}
			fmt.Printf("adding drop at:%v\n", old)
			heights[old]++
		} else {
			fmt.Printf("adding drop at:%v\n", K)
			heights[K]++
		}
	}
	return heights
}
