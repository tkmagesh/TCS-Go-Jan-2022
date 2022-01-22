package main

import (
	"io"
	"os"
	"strings"
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
	stringReader := strings.NewReader("378535 Consequat cillum^(*&(&* adipisicing ea nisi tempor consequat velit commodo nisi dolor culpa magna nulla. Aute in elit ea occaecat elit. Ut laborum dolor occaecat laborum pariatur ipsum adipisicing cupidatat irure cupidatat reprehenderit esse. Amet ut irure Lorem in mollit proident exercitation consectetur magna et cillum aliqua. Eu sit aliqua elit commodo quis amet enim. Id sint et proident nisi aliquip dolore consectetur ullamco ad.")
	alphaReader := AlphaReader{stringReader}
	io.Copy(os.Stdout, alphaReader)
}
