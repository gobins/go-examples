package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("ip4", ":8080")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	fmt.Println(listener)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error: ", err)
		os.Exit(1)
	}
}
