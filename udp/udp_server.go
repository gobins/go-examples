package main

import (
	log "github.com/Sirupsen/logrus"
	"net"
	"os"
)

func main() {
	log.Info("Starting UDP server on port 8080")

	udpAddr, err := net.ResolveUDPAddr("udp", ":8080")
	udpConn, err := net.ListenUDP("udp", udpAddr)

	handleError(err)
	for {
		log.Info("Waiting to serve clients")
		handleClient(*udpConn)
	}
}

func handleClient(conn net.UDPConn) {
	defer conn.Close()
	var buffer [512]byte
	for {
		len, err := conn.Read(buffer[0:])
		handleError(err)
		log.Info(string(buffer[0:len]))
		_, err2 := conn.Write(buffer[0:len])
		handleError(err2)
	}
}

func handleError(err error) {
	if err != nil {
		log.Error("Fatal error: ", err.Error())
		os.Exit(1)
	}
}
