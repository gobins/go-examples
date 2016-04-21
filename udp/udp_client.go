package main

import (
	log "github.com/Sirupsen/logrus"
	"net"
	"os"
)

func main() {
	log.Info("Startng UDP client")

	udpAddr, err := net.ResolveUDPAddr("udp", ":8080")
	handleError(err)

	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	handleError(err)

	for {
		_, err := udpConn.Write([]byte("Ping"))
		handleError(err)
	}
}

func handleError(err error) {
	if err != nil {
		log.Error("Fatal error: ", err.Error())
		os.Exit(1)
	}
}
