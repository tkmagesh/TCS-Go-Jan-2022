package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var opCount int32

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn(wg)
	}
	wg.Wait()
	fmt.Println("opCount = ", opCount)
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt32(&opCount, 1)
	//fmt.Println("fn invoked")
}
