package main

import (
	"fmt"
)

//

func sendDataToChannel(ch chan int, v int) {
	ch <- v
}

type Person struct {
	Name string
	Age  int
}

func sendPerson(ch chan Person, p Person) {
	ch <- p
}

func sendChan(ch chan<- int, in int) {
	ch <- in
}

func sendStringChan(ch chan string, str string) {
	ch <- str
	close(ch)
}

func main() {
	var ci = make(chan int)
	//fmt.Println(ci)
	//var chInt chan int
	//fmt.Println(chInt)  // a channel that is not initialized or zero-value is nil
	go sendDataToChannel(ci, 10)
	v := <-ci
	fmt.Println(v)

	var persons = make(chan Person)

	p := Person{"John", 22}
	go sendPerson(persons, p)
	receivedPerson := (<-persons).Name
	fmt.Println(receivedPerson)

	//	channel can be unidirectional   通道可以是单向的

	var unidirectionalChanSend = make(chan<- int) //
	go sendChan(unidirectionalChanSend, 3)
	//unidirectionalChanSend <- 3
	/*
		ch := make(chan<- data_type)        // The channel operator is after the chan keyword
											// The channel is send-only

		ch := make(<-chan data_type)        // The channel operator is before the chan keyword
										    // The channel is receive-only
	*/
	//var unidirectionalChanReceive = make(<-chan int)
	// <-unidirectionalChanReceive
	fmt.Println("finished")

	var strChan = make(chan string)
	str := "hello"
	go sendStringChan(strChan, str)
	receivedStr, ok := <-strChan
	if !ok {
		fmt.Println("channel closed")
	}
	fmt.Println("received string:", receivedStr)

	_, ok = <-strChan
	if !ok {
		fmt.Println("channel closed")
	}
}
