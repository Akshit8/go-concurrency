package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	asyncExecutorWithSleep()
	fmt.Println("execution time for sleep:", time.Since(now))

	now = time.Now()

	asyncExecutorWithChannels()
	fmt.Println("execution time for channels:", time.Since(now))
}

func asyncExecutorWithChannels() {
	c := make(chan string)

	go task1(c)
	go task2(c)
	go task3(c)
	go task4(c)

	<-c
	<-c
	<-c
	<-c
}

func asyncExecutorWithSleep() {
	go process1()
	go process2()
	go process3()
	go process4()

	// not a good fix
	time.Sleep(1 * time.Second)
}

func task1(c chan string) {
	// below line is blocking code
	time.Sleep(100 * time.Millisecond)
	fmt.Println("process 1 completed")
	c <- "complete"
}

func task2(c chan string) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("process 2 completed")
	c <- "complete"
}

func task3(c chan string) {
	// no wait
	fmt.Println("process 3 completed")
	c <- "complete"
}

func task4(c chan string) {
	time.Sleep(150 * time.Millisecond)
	fmt.Println("process 4 completed")
	c <- "complete"
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

