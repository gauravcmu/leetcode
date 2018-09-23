func climbStairs(n int) int {
	var mapping map[int]int = make(map[int]int, 0)
	return climbStairsHelper(n, &mapping)
}
func climbStairsHelper(n int, mapping *map[int]int) int {
	if n <= 3 {
		return n
	} else {
		if val, ok := (*mapping)[n]; ok {
			return val
		} else {
			//To reach nth stair, either you can take one step from n-1th or 2 from n-2th.
			//ways to reach nth stair = ways to reach n-1 th step & ways to reach n-2th step
			(*mapping)[n] = climbStairsHelper(n-1, mapping) + climbStairsHelper(n-2, mapping)
			return (*mapping)[n]
		}
	}
	return -1
}