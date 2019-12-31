package minimum_window_substring

import (
	"fmt"
	"math"
)

func minWindow(s, t string) string {
	result := make([]byte, 0)
	if len(s) < len(t) {
		return string(result)
	}
	mapping := make(map[byte]int, 0)

	start := 0
	end := 0
	minWindow := math.MaxInt32

	for _, b := range t {
		mapping[byte(b)] = 0
	}

	fmt.Printf("mapping: %v\n", mapping)

	for end = 0; end < len(s); end++ {
		if _, ok := mapping[s[end]]; ok {
			mapping[s[end]] += 1
		} //else it is not in T.

		fmt.Printf("mapping: %v, end:%c\n", mapping, s[end])
		for isSolution(mapping) == true && start <= end {
			if end-start < minWindow {
				minWindow = end - start
				result = []byte(s[start:end+1])
			}

			if _, ok := mapping[s[start]]; ok {
				mapping[s[start]] -= 1
			}
			start++
		}

	}

	return string(result)
}

func isSolution(mapping map[byte]int) bool {
	for _, v := range mapping {
		if v <= 0 {
			fmt.Printf("isSolution() - NO\n")
			return false
		}
	}
	fmt.Printf("isSolution() - YES\n")
	return true
}