package main

import (
	"fmt"
	"time"
)

func compute(label string, value int) {
	for i := 0; i < value; i++ {
		time.Sleep(1*time.Second)
		fmt.Printf("%s: %d\n", label, i)
	}
}

func main() {
	fmt.Println("go concurrency")

	go compute("first", 5)
	go compute("second", 8)

	// holding execution of main routine
	fmt.Scanln()
}