package main

import (
	"fmt"
	"time"
)

func sendValue(label string, c chan string) {
	fmt.Printf("started routine %s\n", label)
	time.Sleep(1*time.Second)
	c <- fmt.Sprintf("Hello World from %s", label)
	fmt.Printf("ended routine %s\n", label)
}

func main() {
	fmt.Println("go channel")

	values := make(chan string, 2)
	defer close(values)

	go sendValue("first", values)
	go sendValue("second", values)

	for i := 0; i < 2; i++ {
		value := <- values
		fmt.Println(value)
	}

	time.Sleep(1*time.Second)
}