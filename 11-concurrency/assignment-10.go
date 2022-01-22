/* Modify the program so that the add function can be executed as a goroutine */
package main

import (
	"fmt"
	"sync"
)

//Don't Commnucate by sharing memory. Instead share memory by communicating

var result int
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	wg.Add(1)
	go add(100, 200)
	wg.Wait()
	mutex.Lock()
	{
		fmt.Println("Result = ", result)
	}
	mutex.Unlock()
}

func add(x, y int) {
	fmt.Printf("processing %d and %d\n", x, y)
	mutex.Lock()
	{
		result = x + y
	}
	mutex.Unlock()
	fmt.Println("returing result : ", result)
	wg.Done()
}
