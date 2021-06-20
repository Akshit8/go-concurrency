package main

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Deposit %d to account with balance %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdraw %d from account with balance %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("go mutex")

	balance = 1000

	var wg sync.WaitGroup
	wg.Add(3)
	go deposit(500, &wg)
	go withdraw(700, &wg)
	go deposit(200, &wg)
	wg.Wait()

	fmt.Printf("net balance: %d\n", balance)
}
