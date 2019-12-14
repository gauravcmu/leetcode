package main

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		i := 0
		j := 0

		indeg := make([]int, n)

		for i = 0; i < n; i++ {
			for j = 0; j< n; j++ {
				if knows(i, j) && i != j{
					indeg[i]--
					indeg[j]++
				}
			}
		}

		for index, val := range indeg {
			if val == n-1 {
				return index
			}
		}

		return -1
	}
}
