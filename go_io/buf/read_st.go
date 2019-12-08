package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	str := "This message contains a few vowels. "

	r := strings.NewReader(str)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Printf("Failed to read %v", err)
	}

	fmt.Printf("%s\n", b)
}
