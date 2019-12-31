package main

import (
"fmt"
)

func main() {
	fmt.Printf("ans: %v \n",checkInclusion("adieb", "eidbaooo"))
}

/*
s1 = "ab"
s2 = "eidbaooo"

find a permutation of ab in s2.

*/

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}

	mapping := make(map[byte]int, 0)
	var i int

	for _, b := range s1 {
		mapping[byte(b)] += 1
	}

	for i = 0; i < len(s1); i++ {
		if _, ok := mapping[s2[i]]; ok {
			mapping[s2[i]] -= 1
		}
	}
	if isSolution(mapping) {
		return true
	}

	for ; i < len(s2); i++ {
		if _, ok := mapping[s2[i]]; ok {
			mapping[s2[i]] -= 1
		}
		if _, ok := mapping[s2[i-len(s1)]]; ok {
			fmt.Printf("set: i:%v %c -1\n",i, s2[i-len(s1)])
			mapping[s2[i-len(s1)]] += 1
		}
		if isSolution(mapping) {
			return true
		}
	}

	return false
}

func isSolution(mapping map[byte]int) bool {
	for k, v := range mapping {
		fmt.Printf("%c:%d\n", k, v)
		if v != 0 {
			return false
		}
	}
	return true
}