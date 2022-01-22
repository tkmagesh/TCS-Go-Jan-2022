package main

import (
	"io"
	"os"
)

type AlphaReader struct {
	Src io.Reader
}

func (alphaReader AlphaReader) Read(p []byte) (int, error) {
	inputData := make([]byte, len(p))
	count, err := alphaReader.Src.Read(inputData)
	if err != nil {
		return count, err
	}
	dataCount := 0
	for idx := 0; idx < len(inputData); idx++ {
		if (inputData[idx] >= 'A' && inputData[idx] <= 'Z') || (inputData[idx] >= 'a' && inputData[idx] <= 'z') {
			p[dataCount] = inputData[idx]
			dataCount++
		}
	}
	return dataCount, io.EOF
}

func main() {
	fileReader, err := os.Open("data1.txt")
	if err != nil {
		panic(err)
	}
	alphaReader := AlphaReader{fileReader}
	io.Copy(os.Stdout, alphaReader)
}
