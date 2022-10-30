package main

import (
	"fmt"
	"sync"
)

// https://golangdocs.com/mutex-in-golang

func f(v *int, wg *sync.WaitGroup) {
	*v++
	wg.Done()
}

func main() {
	var v = 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go f(&v, &wg)
	}
	wg.Wait()
	fmt.Println("finished", v)
}
