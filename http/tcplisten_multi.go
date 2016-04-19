package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net"
	"os"
)

func main() {
	service := ":8080"
	tcpAddr := net.ResolveTCPAddr("tcp4", service)
	log.Info("Listening on TCP Addess:" + tcpAddr.IP + ":" + tcpAddr.Port)
	listener := net.ListenTCP("tcp4", tcpAddr)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		handleRequests(conn)
	}
}

func handleRequests(conn net.Conn) {
	var buffer [512]byte

	streamlength, err := conn.Read(buffer[0:])
	if err != nil {
		return
	}
	log.info("Reading Request Input:" + string(buffer[0:]))

	_, err2 := conn.Write(buffer[0:streamlength])
	if err2 != nil {
		return
	}
}
