package main

import (
	"fmt"
	"sync"
)

type OpCount struct {
	counter int
	sync.Mutex
}

func (opCount *OpCount) Increment() {
	opCount.Lock()
	{
		opCount.counter++
	}
	opCount.Unlock()
}

var opCount OpCount

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn(wg)
	}
	wg.Wait()
	fmt.Println("opCount = ", opCount.counter)
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	opCount.Increment()
	//fmt.Println("fn invoked")
}
