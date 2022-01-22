package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	stop := time.After(20 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Print("Tick")
		case <-stop:
			fmt.Println("Done")
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
