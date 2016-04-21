package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net"
	"os"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":8080")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	log.Info("Starting TCP Listen on :", tcpAddr.Port)
	checkError(err)

	for {
		log.Info("Accepting TCP connections")
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		log.Info("Closing TCP connection")
		conn.Close()
	}
}

// func handleClient(conn net.Conn) {
// 	var buf [512]byte
// 	for {
// 		stream, err := conn.Read(buf[0:])
// 		if err != nil {
// 			return
// 		}
// 		fmt.Println(string(buf[0:]))
// 		_, err2 := conn
// 	}
// }

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error: ", err)
		os.Exit(1)
	}
}
