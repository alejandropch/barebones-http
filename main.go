package main

import (
	"bytes"
	"fmt"
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
	go func() {
		defer close(c)
		defer file.Close()
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
				time.Sleep(500 * time.Millisecond)
				str = ""
			}
			if i != -1 {
				str += string(data[i+1:])
			}
		}
	}()
	return c // 1) return out because remember that the go function will execute on other thread
	// so this will be executed almost immediately

}
func main() {
	file, err := os.Open("./message.txt")
	check(err)

	for i := range getLinesChannel(file) { // 2) it will read from the returned channel when data is piped through it. It will be executed once. But the loop will be woken up when there is a value to be read (from the channel)
		fmt.Printf("%s\n", i)
	}
}
