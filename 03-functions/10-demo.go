package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("deferred from main")
	}()
	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	defer func() {
		fmt.Println("deferred from f1 - [1]")
	}()
	defer func() {
		fmt.Println("deferred from f1 - [2]")
	}()
	defer func() {
		fmt.Println("deferred from f1 - [3]")
	}()
	fmt.Println("f1 started")
	f2()
	fmt.Println("f1 completed")
}

func f2() {
	defer func() {
		fmt.Println("deferred from f2")
	}()
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}
