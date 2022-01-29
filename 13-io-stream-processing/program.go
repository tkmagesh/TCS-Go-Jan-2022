package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataCh := make(chan int)
	oddCh := make(chan int)
	evenCh := make(chan int)
	oddSumCh := make(chan int)
	evenSumCh := make(chan int)

	fileWg := &sync.WaitGroup{}
	processWg := &sync.WaitGroup{}

	fileWg.Add(2)
	go source("data1.dat", dataCh, fileWg)
	go source("data2.dat", dataCh, fileWg)

	processWg.Add(4)
	go splitter(dataCh, oddCh, evenCh, processWg)
	go sum(oddCh, oddSumCh, processWg)
	go sum(evenCh, evenSumCh, processWg)
	go merger(oddSumCh, evenSumCh, "result.txt", processWg)

	fileWg.Wait()
	close(dataCh)

	processWg.Wait()
	fmt.Println("Job Done")
}

func source(fileName string, ch chan int, wg *sync.WaitGroup) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		file.Close()
		wg.Done()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		val, err := strconv.Atoi(str)
		if err == nil {
			ch <- val
		}
	}
}

func splitter(dataCh chan int, oddCh chan int, evenCh chan int, wg *sync.WaitGroup) {
	defer func() {
		close(oddCh)
		close(evenCh)
		wg.Done()
	}()
	for val := range dataCh {
		if val%2 == 0 {
			evenCh <- val
		} else {
			oddCh <- val
		}
	}
}

func sum(numCh chan int, resultCh chan int, wg *sync.WaitGroup) {
	var result int
	for num := range numCh {
		result += num
	}
	resultCh <- result
	wg.Done()
}

func merger(oddSumCh chan int, evenSumCh chan int, resultFile string, wg *sync.WaitGroup) {
	file, err := os.Create(resultFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		file.Close()
		wg.Done()
		close(oddSumCh)
		close(evenSumCh)
	}()

	for i := 0; i < 2; i++ {
		select {
		case oddSum := <-oddSumCh:
			file.WriteString(fmt.Sprintf("Odd Total = %d\n", oddSum))
		case evenSum := <-evenSumCh:
			file.WriteString(fmt.Sprintf("Even Total = %d\n", evenSum))
		}
	}

}
