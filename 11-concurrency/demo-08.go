/* Find the reason for the deadlock */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	resultCh := make(chan int)
	wg.Add(1)
	go add(100, 200, resultCh, wg)
	wg.Wait()
	result := <-resultCh
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
