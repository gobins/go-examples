package main

import (
	"net"
	"time"

	log "github.com/Sirupsen/logrus"
)

func main() {
	log.Info("Listenining on TCP Port 8080")
	listener, err := net.Listen("tcp", ":8080")
	net.Listen(net, laddr)

	if err != nil {
		log.Error("Fatal Error: ", err.Error)
	}
	for {
		log.Info("Handling client request")
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	dateTime := time.Now().String()
	conn.Write([]byte(dateTime))
}
