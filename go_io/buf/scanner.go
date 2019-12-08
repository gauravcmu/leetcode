package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("logs.txt")
	if err != nil {
		fmt.Printf("Failed to open file:%v\n", err)
		return
	}

	r := bufio.NewReader(f)

	s := bufio.NewScanner(r)

	//split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	//	advance, token, err = bufio.ScanWords(data, atEOF)
	//	if err == nil && token != nil {
	//		_, err = strconv.ParseInt(string(token), 10, 32)
	//	}
	//	return
	//}
	s.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			if token[0] == '[' {
				token = token[1:]
			}
			if token[len(token)-1] == ']' {
				token = token[:len(token)-1]
			}
			_, err = strconv.ParseFloat(string(token), 64)
			if err != nil {
				fmt.Printf("err:%v: %s \n", err, token)
			}
		}
		return advance, token, err
	})

	for s.Scan() != false {
		fmt.Printf("%v\n", s.Text())
	}
}
