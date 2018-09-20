func longestValidParentheses(s string) int {
	return helper(s)
}

func helper(s string) int {
	var s1 *stack = &stack{}
	length := 0
	for index, value := range s {
		if value == '(' {
			s1.push(index)
		} else {
			if s1.isEmpty() != true && s[s1.top()] == '(' {
				s1.pop()
				tmp := 0
				if s1.isEmpty() {
					//basically we have a balanced string at this point. index + 1 is the length
					//of this string
					tmp = index + 1
				} else {
					tmp = index - s1.top()
				}
				length = max(length, tmp)
			} else {
				s1.push(index)
			}
		}
	}
	return length
}

type stack []int

func (s *stack) top() int {
	t := (*s)[len(*s)-1]
	return t
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(val int) {
	*s = append((*s), val)
}

func (s *stack) pop() int {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}