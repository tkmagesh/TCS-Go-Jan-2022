package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var opCount int

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
	mutex.Lock()
	{
		opCount++
	}
	mutex.Unlock()
	//fmt.Println("fn invoked")
}
