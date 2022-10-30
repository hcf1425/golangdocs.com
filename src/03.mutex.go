package main

import (
	"fmt"
	"sync"
)

func f2(v *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	*v++
	mu.Unlock()
	wg.Done()

}

func main() {
	var v = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go f2(&v, &wg, &mu)
	}
	wg.Wait()
	fmt.Println("finished:", v)

}
