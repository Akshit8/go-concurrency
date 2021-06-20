package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	process1()
	process2()
	process3()
	process4()

	fmt.Println("execution time:", time.Since(now))
}

func process1() {
	// below line is blocking code
	time.Sleep(100 * time.Millisecond)
	fmt.Println("process 1 completed")
}

func process2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("process 2 completed")
}

func process3() {
	// no wait
	fmt.Println("process 3 completed")
}

func process4() {
	time.Sleep(150 * time.Millisecond)
	fmt.Println("process 4 completed")
}
