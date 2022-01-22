package main

import (
	"fmt"
	"io"
)

type AlphaReader struct {
	str     string
	currPtr int
}

func NewAlphaReader(str string) AlphaReader {
	reader := AlphaReader{str: str}
	reader.currPtr = 0
	return reader
}

//io.Reader interface implementation
func (alphaReader AlphaReader) Read(buffer []byte) (n int, err error) {
	count := 0
	for idx := alphaReader.currPtr; idx < len(alphaReader.str); idx++ {
		if count == len(buffer) {
			alphaReader.currPtr = alphaReader.currPtr + count
			return count, nil
		}
		if (alphaReader.str[idx] >= 'a' && alphaReader.str[idx] <= 'z') || (alphaReader.str[idx] >= 'A' && alphaReader.str[idx] <= 'Z') {
			buffer[count] = alphaReader.str[idx]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	str := "Consequat id eiusmod 23479 laborum non esse consequat cupidatat elit fugiat officia fugiat. Ad culpa aliqua 340 *&^%$&() cupidatat nulla qui minim pariatur cupidatat mollit aliqua velit eu officia. Amet ex voluptate velit dolor nulla nostrud dolor ullamco anim enim. Consequat sit et pariatur esse ipsum amet duis ullamco. Labore ad culpa deserunt est ea. Fugiat duis ad dolore dolor sunt eu qui ipsum aliqua laborum."
	reader := NewAlphaReader(str)
	buffer := make([]byte, 50)
	noOfBytesRead, err := reader.Read(buffer)
	if err == io.EOF {
		println(string(buffer[:noOfBytesRead]))
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(noOfBytesRead)
	fmt.Println(string(buffer[:noOfBytesRead]))
}
