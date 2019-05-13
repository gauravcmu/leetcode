package main

import (
	"fmt"
	"time"
)

/* timer create returns a channel tha blocks the thread until timeout
timer.Stop() stops the timer. Reset() resets it. */
// func main() {
// 	fmt.Printf("Create a timer\n")
// 	done := make(chan bool)

// 	t := time.NewTimer(time.Second * 5)
// 	go timeoutHandler(t.C, done)
// 	//t.Stop()
// 	<-done
// 	fmt.Printf("unblocked\n")

// }

// func timeoutHandler(c <-chan time.Time, done chan bool) {
// 	m := <-c
// 	fmt.Printf("Timeout HIT: %v\n", m)
// 	done <- true
// }

/* Ticker usage */

func main() {
	t := time.NewTicker(time.Second)
	for tt := range t.C {
		fmt.Printf("tick at: %v", tt)
	}

	t.Stop()
}
