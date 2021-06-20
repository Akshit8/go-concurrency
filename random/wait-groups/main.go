package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var urls = []string{
	"https://google.com",
	"https://github.com/Akshit8",
	"https://youtube.com",
}

func fetchStatus(w http.ResponseWriter, req *http.Request) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			fmt.Printf("running routine for %s\n", url)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%v\n", err)
			}
			fmt.Fprintf(w, "%v : %s\n", resp.Status, url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func myFunc(wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	fmt.Println("finished go routine")
	wg.Done()
}

func simpleExample() {
	var wg sync.WaitGroup
	wg.Add(1)
	go myFunc(&wg)
	wg.Wait()
}

func main() {
	fmt.Println("go wait groups")
	// simpleExample()
	// fmt.Println("main.main() finished")

	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":8000", nil))
}