package main

import "fmt"

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		s := &stack{}
		for i := 0; i< n; i++ {
			s.push(i)
		}

		//push them all on stack.
		//Selectively knock off one by one using one comparision.
		//if a knows b -> a cant be celebrity but b potentially is.
		//push b or a accordingly.

		//At the end - have on element left over due eleminations, this may
		//maynot be the celeb. Do one round of validation before declaring.
		for len(*s) > 1 {
			a := s.pop()
			b := s.pop()

			if knows(a, b) {
				s.push(b)
			} else {
				s.push(a)
			}
		}

		fmt.Printf("length: %v\n", len(*s))
		//at this point we expect only one element in stack.
		c := s.pop()



		//double check if c is the celebrity.
		for i:=0; i< n;i++ {
			if i != c && (knows(c, i) || !knows(i, c)) {
				return -1
			}
		}
		return c
	}
}


type stack []int

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(v int) {
	(*s) = append(*s, v)
}

func (s *stack) pop() int {

	t := (*s)[len(*s)-1]
	// fmt.Printf("popped:%v len:%v\n", t, len(*s))
	(*s) = (*s)[:len(*s)-1]
	return t
}
