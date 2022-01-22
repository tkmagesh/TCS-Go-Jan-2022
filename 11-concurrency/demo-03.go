package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	WaitGroup Methods
		- Add() => increment the counter
		- Done() => decrement the counter
		- Wait() => block the execution of the function until the counter becomes zero
*/
func main() {
	//var wg sync.WaitGroup
	wg := &sync.WaitGroup{}
	fmt.Println("Main started")

	wg.Add(1)
	go f1(wg)

	wg.Add(1)
	go f2(wg)

	wg.Wait()
	fmt.Println("Main completed")
}

func f1(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("f1 is invoked")
	wg.Done()
}

func f2(wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	fmt.Println("f2 is invoked")
	wg.Add(1)
	f3(wg)
	wg.Done()
}

func f3(wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	fmt.Println("f3 is invoked")
	wg.Done()
}
