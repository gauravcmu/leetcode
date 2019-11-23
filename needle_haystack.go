package main

import (
	"fmt"
)

func main() {
	haystack := "aabcdeflmnaaaabcpqr"
	needle := "abc"
	
	fmt.Printf("%v", find(haystack, needle))
	
	fmt.Printf("%v", find("abcd", "zy"))
	
	fmt.Printf("%v", find("abcdabcd", "abcd"))
}

func find(haystack, needle string) []int {
	result := make([]int, 0)
	
	for i:=0; i < len(haystack);i++ {
		for j:=0; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				break
			} else if j == len(needle)-1 {
				//found a solution. 
				result = append(result, i)
			}
		}
	}
	return result 

}
