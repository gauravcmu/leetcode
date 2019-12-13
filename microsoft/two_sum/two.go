package main

import (
	"fmt"
)

func main() {

}

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int, 0) //map integer to index.
	for index, value := range nums {
		if _, ok := mapping[target - value]; !ok {
			mapping[target - value] = index
		}
	}

	for index, value := range nums {
		if i, ok := mapping[target-value]; ok {
			return []int{index, i}
		}
	}
	return []int{}
}