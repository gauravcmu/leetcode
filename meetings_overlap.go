package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type meeting struct {
	start int
	end   int
}

func readInput() ([]meeting, error) {
	meetings_list := make([]meeting, 0)

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Errorf("Failed to read from stdin")
		return meetings_list, err
	}
	//  fmt.Printf("input: %s\n",input)

	input2 := strings.Split(string(input), "\n")
	if len(input2)%2 == 0 {
		fmt.Errorf("invalid Meeting input")
		os.Exit(1)
	}

	// create a slice of meetings represented by types meeting with start and end times.
	for i := 0; i < len(input2); i++ {
		//create a meeting - parse input to get integer start / end times for the meeting
		meet := meeting{}
		//input2 has meetings as each element in slice - split and create slice of start_end
		start_end := strings.Fields(input2[i])

		//    fmt.Printf("start_end %v", start_end)

		meet.start, err = strconv.Atoi(string(start_end[0]))
		meet.end, err = strconv.Atoi(string(start_end[1]))

		//append this to meetings list - which will be used for the overlap logic.
		meetings_list = append(meetings_list, meet)
		//    fmt.Printf("Meetings list is:%v", meetings_list)
	}

	return meetings_list, err
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	meetings_list, err := readInput()
	if err != nil {
		fmt.Errorf("Failed to read from stdin")
		return
	}

	//got input in meetings_list - slice of meeting type.

	//go over the meetings starting from second meeting in the input.
	//start_first_meeting, end_first_meeting

	start_first_meeting := meetings_list[0].start
	end_first_meeting := meetings_list[0].end

	for i := 1; i < len(meetings_list); i++ {
		//if first meeting starts after this meeting starts but starts before it ends - yes.
		if start_first_meeting >= meetings_list[i].start && start_first_meeting < meetings_list[i].end {
			fmt.Printf("yes\n")
			continue
		}
		//if first meeting ends before the next meeting but after next meeting starts - yes
		if end_first_meeting > meetings_list[i].start && end_first_meeting <= meetings_list[i].end {
			fmt.Printf("yes\n")
			continue
		}
		fmt.Printf("no\n")
	}
}
