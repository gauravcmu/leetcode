package longest_substring_no_repetition

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	mapping := make(map[byte]int, 0)

	start := 0
	end := 0

	max := 0

	for end = 0; end < len(s); end++ {
		mapping[s[end]] += 1

		for start < len(s) && start <= end && isSolution(mapping) != true {
			if end-start > max {
				max = end - start
			}
			mapping[s[start]] -= 1
			if mapping[s[start]] == 0 {
				delete(mapping, s[start])
			}
			start++
		}
	}
	return max
}

func isSolution(mapping map[byte]int) bool {
	for _, v := range mapping {
		if v > 1 {
			return false
		}
	}

	return true
}
