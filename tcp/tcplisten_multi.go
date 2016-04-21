package main

import (
	log "github.com/Sirupsen/logrus"
	"net"
	"os"
)

func main() {
	service := ":8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	log.Info("Listening on TCP Port:", tcpAddr.Port)
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		go handleRequests(conn)
	}
}

func handleRequests(conn net.Conn) {
	defer conn.Close()
	var buffer [512]byte

	streamlength, err := conn.Read(buffer[0:])
	if err != nil {
		return
	}
	log.Info("Reading Request Input:" + string(buffer[0:]))

	_, err2 := conn.Write(buffer[0:streamlength])
	if err2 != nil {
		return
	}
}

func checkError(err error) {
	if err != nil {
		log.Error("Fatal error: " + err.Error())
		os.Exit(1)
	}
}
