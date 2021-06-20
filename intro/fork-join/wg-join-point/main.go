package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)

	// work
	go func() {
		defer wg.Done()
		work()
	}()
	// main work
	time.Sleep(100 * time.Millisecond)

	// kind of join point
	wg.Wait()

	fmt.Println("main finsished:", time.Since(now))
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("work done")
}
