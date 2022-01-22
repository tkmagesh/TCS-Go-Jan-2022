/* Modify the program so that the add function can be executed as a goroutine */
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
	go func() {
		resultCh <- add(100, 200)
		wg.Done()
	}()
	result := <-resultCh
	wg.Wait()
	fmt.Println("Result = ", result)

}

func add(x, y int) int {
	fmt.Printf("processing %d and %d\n", x, y)
	time.Sleep(2 * time.Second)
	result := x + y
	fmt.Println("returing result : ", result)
	return result
}
