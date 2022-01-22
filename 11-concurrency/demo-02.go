package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

/*
	WaitGroup Methods
		- Add() => increment the counter
		- Done() => decrement the counter
		- Wait() => block the execution of the function until the counter becomes zero
*/
func main() {
	fmt.Println("Main started")

	wg.Add(1)
	go f1()

	wg.Add(1)
	go f2()

	wg.Wait()
	fmt.Println("Main completed")
}

func f1() {
	time.Sleep(2 * time.Second)
	fmt.Println("f1 is invoked")
	wg.Done()
}

func f2() {
	time.Sleep(5 * time.Second)
	fmt.Println("f2 is invoked")
	wg.Done()
}
