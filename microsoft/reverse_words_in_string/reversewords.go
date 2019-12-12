package main

import (
	"fmt"
	"strings"
)

func main() {
	reverseWords("the sky is blue")

}

func reverseWords(s string) string {
	tokens := strings.Fields(s)

	for i, j := 0, len(tokens) -1; i<j; i,j = i+1, j-1 {
		tokens[i], tokens[j] = tokens[j], tokens[i]
	}

	result := make([]byte, 0)

	for index, token := range tokens {
		result = append(result, []byte(token)...)
		if index != len(tokens)-1 {
			result = append(result, []byte(" ")...)
		}
	}
	return string(result)
}