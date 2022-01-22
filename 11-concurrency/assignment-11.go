package main

import (
	"fmt"
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
	print("Hello")
	print("World")
}

func print(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}
