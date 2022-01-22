package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var txt = scanner.Text()
		if txt == "exit" {
			break
		}
		fmt.Println(txt)
	}
}
