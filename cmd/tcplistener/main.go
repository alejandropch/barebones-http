package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal("error", e)
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
		if str != "" {
			// finding a possible \n in the lasting str variable
			if i = bytes.IndexByte([]byte(str), '\n'); i != -1 {
				c <- str[:i]
			}
		}
	}()
	return c // 1) return the string channel because remember that the go function will be executed on other thread
	// so this will be executed almost immediately

}
func main() {
	listener, err := net.Listen("tcp", ":1234")
	check(err)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		check(err)
		for i := range getLinesChannel(conn) { // 2) it will read from the returned channel when data is piped through it. It will be executed once. But the loop will be woken up when there is a value to be read (from the channel)
			fmt.Printf("%s\n", i)
		}
		conn.Close()
	}
}
