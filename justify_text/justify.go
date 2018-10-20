import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	length := 0
	r := []byte{}
	res := make([]string, 0)

	for _, word := range words {
		if length+1+len(word) <= maxWidth {
			length += len(word)
			length += 1

			r = append(r, []byte(word)...)
			r = append(r, []byte(" ")...)
		} else {
			res = append(res, justify(string(r), maxWidth))
			//res = append(resdd, string('\n'))
			r = []byte{}
			length = 0

			r = append(r, []byte(word)...)
			r = append(r, []byte(" ")...)
			length += len(word)
			length += 1
		}
	}
	if len(r) != 0 {
		fmt.Printf("-- %v:r len(r):%v \n", string(r), len(r))
		l := maxWidth - len(r)
		for i := 0; i < l; i++ {
			r = append(r, ' ')
			fmt.Printf("=== len of r == %v max %v i:%v\n", len(r), maxWidth, i)
		}
		res = append(res, string(r))
	}

	return res
}

func justify(input string, maxWidth int) string {
	input = strings.TrimSuffix(input, " ")
	li := len(input)

	diff := maxWidth - li

	res := []byte{}

	count := 0
	for _, val := range input {
		if val == ' ' {
			count++
		}
	}

	factor := 0
	if count > 0 {
		factor = diff / count
	}

	for _, val := range input {
		res = append(res, byte(val))
		if val == ' ' && diff > 0 {
			fmt.Printf("Found space. \n")
			for i := 0; i < factor; i++ {
				res = append(res, byte(val))
				diff--
			}
			if factor == 0 && diff > 0 {
				res = append(res, byte(val))
				diff--
			}
		}
	}
	for diff != 0 {
		res = append(res, ' ')
		diff--
	}
	fmt.Printf("input:[%v] res:[%v]\n", input, string(res))
	return string(res)
}