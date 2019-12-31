package longest_substring_at_most_k_distinct_chars

import "math"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if len(s) < k || s == ""{
		return len(s)
	}
	if k == 0 {
		return 0
	}

	start := 0	
	end := 0
	max := math.MinInt32

	mapping := make(map[byte]int, 0)

	for end = 0; end < len(s); end++ {
		mapping[s[end]] += 1

		if isSolution(mapping, k) == true {
			if getLen(mapping) > max {
				max = getLen(mapping)
			}
		} else {
			for ; start < len(s) && isSolution(mapping, k) != true; {
				mapping[s[start]] -= 1
				if mapping[s[start]] == 0 {
					delete(mapping, s[start])
				}
				start++
			}
		}
	}
	return max
}

func isSolution(mapping map[byte]int, k int) bool {
	return len(mapping) <= k
}

func getLen(mapping map[byte]int) int {
	count := 0
	for _, v := range mapping {
		count += v
	}

	return count
}
