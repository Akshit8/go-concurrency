package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	concurrencyWithWaitGroups()
}

func concurrencyWithWaitGroups() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func () {
		count("first")
		wg.Done()
	}()

	go func () {
		count("second")
		wg.Done()
	}()

	wg.Wait()
}

func count(str string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: %s\n", i, str)
		time.Sleep(500*time.Millisecond)
	}
}