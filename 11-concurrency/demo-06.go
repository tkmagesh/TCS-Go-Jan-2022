/* Modify the program so that the add function can be executed as a goroutine */
package main

import (
	"fmt"
	"sync"
	"time"
)

/*
var wg sync.WaitGroup
var resultCh chan int = make(chan int)

func main() {
	wg.Add(1)
	go add(100, 200)
	result := <-resultCh
	wg.Wait()
	fmt.Println("Result = ", result)

}

func add(x, y int) {
	fmt.Printf("processing %d and %d\n", x, y)
	time.Sleep(2 * time.Second)
	result := x + y
	fmt.Println("returing result : ", result)
	resultCh <- result
	wg.Done()
}
*/

func main() {
	wg := &sync.WaitGroup{}
	resultCh := make(chan int)
	wg.Add(1)
	go add(100, 200, resultCh, wg)
	result := <-resultCh
	wg.Wait()
	fmt.Println("Result = ", result)

}

func add(x, y int, resultCh chan int, wg *sync.WaitGroup) {
	fmt.Printf("processing %d and %d\n", x, y)
	time.Sleep(2 * time.Second)
	result := x + y
	fmt.Println("returing result : ", result)
	resultCh <- result
	wg.Done()
}
