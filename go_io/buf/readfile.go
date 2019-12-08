package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
)

func main() {
	//read a file with buffered io.
	f, err := os.Open("readfile.go")
	if err != nil {
		fmt.Printf("failed to open file:%v\n", err)
		return
	}

	r := bufio.NewReader(f)


	p := make([]byte, 10)
	n := 1
	for n != 0 {
		n, err := r.Read(p)
		if err != nil && err != io.EOF {
			fmt.Printf("failed to readr:%v\n", err)
			return
		}
		if n == 0 {
			break
		}
		fmt.Printf("%s", p)
	}

}
