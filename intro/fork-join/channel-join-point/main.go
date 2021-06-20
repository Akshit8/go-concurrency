package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	c := make(chan struct{})

	// work
	go func() {
		work()
		c <- struct{}{}
	}()
	// main work
	time.Sleep(100 * time.Millisecond)

	// this is join point
	// will wait until any routines passes
	<-c
	
	fmt.Println("main finsished:", time.Since(now))
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("work done")
}
