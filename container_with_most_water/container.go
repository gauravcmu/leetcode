func maxArea(height []int) int {
    max_so_far := 0 
    for i, j := 0, len(height) -1; i<j; {
        area := (j - i) * min (height[j], height[i])
        if area > max_so_far {
            max_so_far = area
        }
        
        //move shorter guy to hopefully get the next one taller.
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
    }
    return max_so_far
}

func min(a, b int) int {
    if a < b {
        return a 
    }
    return b
}
