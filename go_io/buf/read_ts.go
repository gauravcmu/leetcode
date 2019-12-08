package main

import (
    "flag"
    "fmt"
   "io/ioutil"
)

func main() {
      filename := flag.String("filename", "test.txt", "pass file name")
      flag.Parse()

      b, err := ioutil.ReadFile(*filename)
      if err != nil {
          fmt.Printf("Failed to read %v", err)
      }

      fmt.Printf("%s\n", b)
}


