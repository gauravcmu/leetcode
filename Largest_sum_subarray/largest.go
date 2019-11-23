package main

import (
	"fmt"
)

func main() {
	input := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	fmt.Printf("%v\n", findMaxSubArray(input))
}

//kadanes algorithm 
func findMaxSubArray(input []int) int {
	max_sum := 0
	max_so_far := 0

	for i := 0; i < len(input); i++ {
		max_so_far += input[i]
		if max_so_far < 0 {
			//negative - reset to 0
			max_so_far = 0
		} else if max_so_far > max_sum {
			max_sum = max_so_far
		}
	}
	return max_sum 
}

