package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	lAddr, err := net.ResolveUDPAddr("udp", ":1234")
	check(err)
	rAddr, err := net.ResolveUDPAddr("udp", ":4321")
	check(err)
	conn, err := net.DialUDP("udp", lAddr, rAddr)
	check(err)
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">")
		str, err := reader.ReadString('\n')
		check(err)
		_, err = conn.Write([]byte(str))
		check(err)
	}

}
