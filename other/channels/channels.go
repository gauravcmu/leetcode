package main

import (
	"fmt"
	"time"
)

/* 1 */
func foo(done chan bool) {
	for i := 0; i < 1000; i++ {
		fmt.Printf("%v inside foo\n", i)
	}
	(done) <- true
}

// func main() {
// 	done := make(chan bool)

// 	fmt.Printf("Calling go routine foo\n")
// 	go foo(done)

// 	<-done
// 	fmt.Printf("Done from go routine foo\n")

// }

/* 2 */
// func writeToChan(c chan string) {
// 	for i := 0; i < 1000; i++ {
// 		fmt.Printf("Writing:\n")
// 		c <- time.Now().String()
// 	}
// }

// func readFromChan(c chan string, done chan bool) {
// 	for i := 0; i < 1000; i++ {
// 		fmt.Printf("Read:%v\n", <-c)
// 	}

// 	done <- true
// }
// func main() {
// 	c := make(chan string, 10)
// 	done := make(chan bool)

// 	go writeToChan(c)
// 	go readFromChan(c, done)

// 	<-done
// }

/*3 range and close - range iterates untill channnel is closed */

// func iterate(c chan string) {
// 	for i := 0; i < 100; i++ {
// 		fmt.Printf("Writing to chan\n")
// 		c <- "hello" + string(i)
// 	}
// 	close(c) //without close it deadlocks on range.
// }

// func main() {
// 	c := make(chan string)
// 	go iterate(c)

// 	for val := range c {
// 		fmt.Printf("%v\n", val)
// 	}
// }

/* 4 use select case on multiple channels */

func makeMeAChannel() chan string {
	c := make(chan string)
	return c
}

func writer(c chan string) {
	for i := 0; i < 1000; i++ {
		fmt.Printf("Writing to:%v\n", c)
		c <- time.Now().String()
	}
}

func main() {
	c1 := makeMeAChannel()
	go writer(c1)
	c2 := makeMeAChannel()
	go writer(c2)

	c3 := makeMeAChannel()
	go writer(c3)

	for i := 0; i < 9; i++ {
		select {
		case v := <-c1:
			fmt.Printf("from C1:%v\n", v)
		case v := <-c2:
			fmt.Printf("from C2:%v\n", v)
		case v := <-c3:
			fmt.Printf("from C3:%v\n", v)
		}
	}
}
