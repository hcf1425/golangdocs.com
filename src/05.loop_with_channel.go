package main

import "fmt"

// https://golangdocs.com/channels-in-golang

func multiplyChan(ch chan int, v int) {
	ch <- v
	ch <- v * 2
	ch <- v * 3
	ch <- v * 7
	ch <- v * 11
	close(ch)
}

func main() {
	var ch = make(chan int)
	var v = 2
	go multiplyChan(ch, v)

	for value := range ch {
		fmt.Println(value)
	}
	fmt.Println("process finished!")

}
