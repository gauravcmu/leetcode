// Given two strings s and t, determine if they are isomorphic.

// Two strings are isomorphic if the characters in s can be replaced to get t.

// All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

// Example 1:

// Input: s = "egg", t = "add"
// Output: true
// Example 2:

// Input: s = "foo", t = "bar"
// Output: false
// Example 3:

// Input: s = "paper", t = "title"
// Output: true
// Note:
// You may assume both s and t have the same length.

func isIsomorphic(s string, t string) bool {
	mapping := make(map[rune]rune, 0)

	for index, c := range s {
		if cmap, ok := mapping[c]; ok {
			//found this char in mapping.
			if index < len(t) && cmap == rune(t[index]) {
				continue
			} else {
				//found a char for which mapping exists and does not look like the same.
				return false
			}
		} else {
			//mapping does not exist for this char.
			for _, value := range mapping {
				if value == rune(t[index]) {
					return false
				}
			}
			mapping[c] = rune(t[index])
		}
	}
	return true
}