package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getLinesChannel(file io.ReadCloser) <-chan string {

	var i int
	str := ""
	c := make(chan string)
	for {
		data := make([]byte, 8)
		counter, err := file.Read(data)
		if err != nil {
			break
		}

		if i = bytes.IndexByte(data, '\n'); i == -1 {
			str += string(data[:counter])
		} else {
			str += string(data[:i])
			c <- str
			time.Sleep(1 * time.Second)
			str = ""
		}
		if i != -1 {
			str += string(data[i+1:])
		}
	}
	return c
}
func main() {
	_, err := os.Open("./message.txt")
	check(err)
	//	for i range getLinesChannel(file)

}
