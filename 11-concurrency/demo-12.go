/* Using range construct on channels */
package main

import (
	"fmt"
	"time"
)

func main() {
	noCh := make(chan int)
	go genFibonacci(noCh)
	for no := range noCh {
		fmt.Println(no)
	}
}

func genFibonacci(resultCh chan int) {
	x, y := 0, 1
	for {
		resultCh <- x
		time.Sleep(500 * time.Millisecond)
		x, y = y, x+y
	}
}
