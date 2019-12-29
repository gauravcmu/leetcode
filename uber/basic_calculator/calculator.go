package basic_calculator

import "fmt"

/*
	- Say we have a sample input (a + b - (c + d))
	- we actually want to calculate a+b and then subtract c+d
	- subtraction can be treated as addition of value with negative sign.
	- so we build result by adding value of number * sign to result.

	- Overall algorithm:
	- Keep calculating the result iteratively.
		- when you see a digit - find the biggest number that is formed with this digit.
		- so starting from i+1 till end of input or you encounter non-digit, keep building number.
		- add the number to result with sign.
   - if character is a non-digit:
		- If this is + or - just update the sign to 1 or -1 accordingly.
		- if it is a space - ignore.

        - if it is an opening bracket, push the current result computed so far to stack. also push sign.
		- ^ stack will have sign of computation so far and the result so far on the stack.
		- eg. [+, a+b]
		- result result variable to 0 and sign to positive for the bracket calculation.

        - if closing bracket - to the result calculated in this bracket, add the result from previous step.
		- so pop the stack to get sign and the result from prev step. add this to current result.
    - at the end of the input - return the result.
 */

func calculate(s string) int {
	st := &stack{}
	result := 0
	sign := 1 //1 for + and -1 for -
	num := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if isDigit(s[i]) == true {
			// 123 + 22
			num = int(ch - '0')
			for ; i+1 < len(s) && isDigit(s[i+1]); {
				num = num*10 + int(s[i+1]-'0')
				i++
			}
			result += num * sign
		} else if ch == '+' {
			sign = 1
		} else if ch == '-' {
			sign = -1
		} else if ch == '(' {
			st.push(result)
			st.push(sign)

			result = 0
			sign = 1
		} else if ch == ')' {
			// ( 12 - (13 + 14))
			// result = -27 + 12

			sg := st.pop()
			v := st.pop()
			fmt.Printf("st: %+v %v %v rs:%v\n", st, sg, v, result)
			result = sg*result + v
		}
	}

	return result
}

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

type stack []int

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(n int) {
	(*s) = append(*s, n)
}

func (s *stack) pop() int {
	t := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return t
}

