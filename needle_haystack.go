package main

import "fmt"

func main() {
	fmt.Printf("result: %v", grep("aaabcdddbbddddabcdefghi", "abc"))
}

func grep(haystack, needle string) []int {
	i := 0
	j := 0

	start_index := 0
	result := make([]int, 0)
	count_match := 0

	for i, j = start_index, 0; i < len(haystack) && j < len(needle); {
		if needle[j] == haystack[i] {
			count_match++
			if count_match == len(needle) {
				fmt.Printf("Add to result %v\n", start_index)
				result = append(result, start_index)
				i++
				j = 0
				start_index = i
				count_match = 0
			} else {
				// match
				i++
				j++
			}
		} else {
			if count_match == 0 {
				if i+1 == len(haystack) {
					return result
				} else {
					i++
				}
			}
			start_index = i
			j = 0
			count_match = 0
		}
	}

	return result
}

