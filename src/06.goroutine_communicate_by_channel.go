package main

import (
	"fmt"
	"time"
)

// not
// https://golangdocs.com/multiple-goroutines-in-golang

func writeChan(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func readChan(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	var ch = make(chan int)
	go readChan(ch)
	go writeChan(ch)
	time.Sleep(time.Second)
	fmt.Println("the process finished")
	//	 as it can be seen channels are the way multiple goroutines communicate.
}
