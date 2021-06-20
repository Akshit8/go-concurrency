package main

import (
	"fmt"
	"time"
)

func main() {
	go work() // fork point

	time.Sleep(100 * time.Millisecond)

	fmt.Println("main finsished")

	// join point for above forked process
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("work done")
}