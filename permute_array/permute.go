package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", permute([]int{1,2,3}))
	
}

func permute(input []int) [][]int {
	result := make([][]int, 0)
	
	if len(input) == 0 {
		return result 
	}
	
	permuteHelper(input, []int{}, &result, "  ")
	return result 
}

func permuteHelper(input []int, chosen []int, result *[][]int, path string) {
	fmt.Printf("%vinput:%v chosen:%v result:%v\n", path, input, chosen, result)
	if len(input) == 0 {
		*result = append(*result, chosen) 
		return 
	} else {
		for i:=0; i< len(input); i++ {
			chosen := append(chosen, input[i])
			permuteHelper(append(append([]int{}, input[:i]...), input[i+1:]...), chosen, result, path+"  ") 
			chosen = chosen[:len(chosen)-1]
		}
	}
}

/*
permuteHelper([1, 2, 3], [], [])
	permuteHelper([2, 3], [1], [])
		permuteHelper([3], [1, 2], [])
			permuteHelper([], [1, 2, 3], [])
		permuteHelper([2], [1, 3], [])
			permuteHelper([], [1, 3, 2], [])
			
		
	permuteHelper([1, 3], [2], [])
		permuteHelper([3], [2, 1], [])
			permuteHelper([], [2, 1, 3], [])
		permuteHelper([1], [2, 3], [])
			permuteHelper([], [2, 3, 1], [])
			
	permuteHelper([1, 2], [3], [])
		permuteHelper([2], [3, 1], [])
			permuteHelper([], [3, 1, 2], [])
		permuteHelper([1], [3, 2], [])
			permuteHelper([], [3, 2, 1], [])
*/
