// Write a function that takes a string as input and returns the string reversed.

// Example 1:

// Input: "hello"
// Output: "olleh"
// Example 2:

// Input: "A man, a plan, a canal: Panama"
// Output: "amanaP :lanac a ,nalp a ,nam A"

func reverseString(s string) string {
	i := 0
	j := len(s) - 1

	if len(s) <= 1 {
		return s
	}
	s2 := []byte(s) //s[i] cannot be written - so use byte array
	for ; i < j; i, j = i+1, j-1 {
		s2[i], s2[j] = s2[j], s2[i]
	}
	return string(s2)
}