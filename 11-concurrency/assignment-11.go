package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Expected outcome:
		Hello
		World
		Hello
		World
		Hello
		World
		Hello
		World
		Hello
		World

	IMPORTANT:
		The loop should remain in the print function
*/
func main() {
	x := make(chan bool)
	y := make(chan bool)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go print("Hello", x, y, wg)
	go print("World", y, x, wg)
	x <- true
	wg.Wait()
}

func print(s string, in, out chan bool, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		<-in
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
		out <- true
	}
	fmt.Println(s, " - Done")
	wg.Done()
}
