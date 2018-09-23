func climbStairs(n int) int {
	var solution []int = make([]int, n+1)
	if n <= 3 {
		return n
	}
	solution[0] = 0
	solution[1] = 1
	solution[2] = 2
	for i := 3; i <= n; i++ {
		solution[i] = solution[i-1] + solution[i-2]
	}
	return solution[n]
}