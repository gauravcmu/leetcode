func singleNumber(nums []int) int {
	sum := make([]int, 64)
    fmt.Println("arch", runtime.GOARCH)
    fmt.Println("int", unsafe.Sizeof(int(0)))

	for _, v := range nums {
		for i := 0; i < 64; i++ {
			if sum[i] != 0 {
				sum[i] += 1 & (v >> uint(i))
			} else {
				sum[i] = 1 & (v >> uint(i))
			}
		}
	}

    res := 0 
	for k, v := range sum {
		if (v%3) != 0 {
            res |= int((v%3) << uint(k)) 
		}
	}
    fmt.Printf("res %+v\n", res) // prints -4 
    return res 
}
