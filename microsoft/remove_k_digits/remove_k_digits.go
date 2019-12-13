package main
func removeKdigits(num string, k int) string {
	s := &stack{}

	if len(num) == 0  || len(num) <= k  {
		return "0"
	}

	s.Push(num[0])
	for i := 1; i< len(num); i++   {
		for ; s.isEmpty()!=true && k > 0 && num[i] < s.Peek(); k-- {
			s.Pop()
		}
		s.Push(num[i])
	}

	//pop k elements for special case strings like 1111111..
	for ;k>0 && s.isEmpty()!=true; k-- {
		s.Pop()
	}

	//remove leading zeros - can't pop..
	for ;s.isEmpty() != true && (*s)[0] == '0'; {
		*s = (*s)[1:]
	}

	//consider the case "10", 1 -> will end up being empty string..
	if len(*s) == 0 {
		return "0"
	}
	return string(*s)
}


type stack []byte

func (s *stack) isEmpty() bool {
	return len((*s)) == 0
}

func (s *stack) Push(v byte) {
	*s = append(*s, v)
}

func (s *stack) Pop() byte {
	t := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return t
}

func (s *stack) Peek() byte {
	return (*s)[len(*s)-1]
}
