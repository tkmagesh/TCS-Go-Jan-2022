package main

import (
	"fmt"
	"time"
)

func main() {
	ch := getEvenNos()
	for i := 0; i < 20; i++ {
		evenNo := <-ch
		fmt.Println("Even No : ", evenNo)
	}

}

func getEvenNos() <-chan int {
	evenCh := make(chan int)
	var evenNo = 0
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			evenNo += 2
			evenCh <- evenNo
		}
	}()
	return evenCh
}
