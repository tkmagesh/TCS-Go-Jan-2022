package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileReader, err := os.Open("data2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fileReader.Close()
	scanner := bufio.NewScanner(fileReader)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		fmt.Printf("line #%d : %s\n", lineCount, scanner.Text())
	}
}
