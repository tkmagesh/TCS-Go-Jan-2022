/* Using range construct on channels */
package main

import (
	"fmt"
	"time"
)

func main() {
	noCh := make(chan int)
	doneCh := make(chan string)
	go genFibonacci(noCh, doneCh)
	go func() {
		fmt.Println("hit ENTER to exit")
		var input string
		fmt.Scanln(&input)
		doneCh <- "stop"
	}()
	for no := range noCh {
		fmt.Println(no)
	}
}

func genFibonacci(resultCh chan int, doneCh chan string) {
	x, y := 0, 1
	for {
		select {
		case <-doneCh:
			close(resultCh)
			return
		case resultCh <- x:
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		}
	}
}
