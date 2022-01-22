/* Find the reason for the deadlock */
package main

import (
	"fmt"
	"time"
)

func main() {

	resultCh := make(chan int)
	go add(100, 200, resultCh)
	result := <-resultCh
	fmt.Println("Result = ", result)

}

func add(x, y int, resultCh chan int) {
	fmt.Printf("processing %d and %d\n", x, y)
	time.Sleep(2 * time.Second)
	result := x + y
	fmt.Println("returing result : ", result)
	resultCh <- result

}
