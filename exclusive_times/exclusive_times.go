import (
	"fmt"
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	result := make([]int, n)
	var tempStack *stack = &stack{}

	for _, v := range logs {
		i := strings.Split(v, ":")
		fname, _ := strconv.Atoi(i[0])
		time, _ := strconv.Atoi(i[2])
		loginfo := info{
			fname:    fname,
			is_start: i[1] == "start",
			time:     time,
		}

		if loginfo.is_start == true {
			//push on the stack
			tempStack.push(loginfo)
		} else {
			fmt.Printf("end:%+v\n", loginfo)
			//end - so pop the top
			top := tempStack.pop()
			fmt.Printf("result:%v\n", result)
			result[loginfo.fname] += loginfo.time - top.time + 1
			if tempStack.isEmpty() != true {
				temp := tempStack.peek()
				result[temp.fname] -= loginfo.time - top.time + 1
			}
		}

	}
	return result
}

type info struct {
	fname    int
	is_start bool
	time     int
}

type stack []info

func (s *stack) push(val info) {
	fmt.Printf("push: %+v\n", val)
	*s = append(*s, val)
}

func (s *stack) pop() info {
	t := (*s)[len(*s)-1]
	fmt.Printf("pop: %v\n", t)
	*s = (*s)[:len(*s)-1]
	return t
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) peek() info {
	fmt.Printf("peek: %v\n", (*s)[len(*s)-1])
	return (*s)[len(*s)-1]
}