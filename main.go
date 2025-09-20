package main

import (
	"bytes"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./message.txt")
	check(err)
	str := ""
	var i int
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
			fmt.Printf("gAA: %s\ni:%d\n", str, i)
			str = ""
		}
		if i != -1 {
			str += string(data[i+1:])
		}
	}

}
