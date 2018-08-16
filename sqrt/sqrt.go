func mySqrt(x int) int {
	return sqrt(x, 1, x-1)
}
func sqrt(x int, start, end int) int {
	if x == 0 || x == 1 {
		return x
	}

    if x == 2 {
        return 1
    }
	if start > end {
		return 0
	}
	mid := (start + end) / 2
	//^ replace with start + (end-start)/2

	temp := mid * mid
	temp2 := (mid - 1) * (mid - 1)
	//temp3 := (mid +1) * (mid+1)
	if temp == x {
		return mid
	} else if temp > x {
		if temp2 < x {
			return mid - 1
		} else {
			return sqrt(x, start, mid-1)
		}
	} else {
		return sqrt(x, mid+1, end)
	}
}
