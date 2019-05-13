package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("fail")
		return
	}

	fmt.Printf("%v", string(input))
}
