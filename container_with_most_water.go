func maxArea(height []int) int {
    maxarea := 0 
    for i:=0; i<len(height)-1; i++ {
        for j:=len(height)-1;j>i;j-- {
            if (j-i) * min(height[j], height[i]) > maxarea {
                fmt.Printf("%v %v\n", i,j)
                maxarea = (j-i) * min(height[j], height[i]) 
            }
        }
    }
    return maxarea
}

func min(a,b int) int {
    if a<b {
        return a
    }
    return b
}
