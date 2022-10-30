package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// https://golangdocs.com/atomic-operations-in-golang-atomic-package
func calAdd(v *uint32, wg *sync.WaitGroup) {
	for i := 0; i < 3000; i++ {
		//time.Sleep(time.Millisecond)
		atomic.AddUint32(v, 2)
		//*v++
	}
	wg.Done()
}

func main() {
	var v uint32 = 0
	var wg sync.WaitGroup
	wg.Add(2)
	go calAdd(&v, &wg)
	go calAdd(&v, &wg)
	wg.Wait()
	fmt.Println("value of v:", v)
}
