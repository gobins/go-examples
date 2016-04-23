package main

import (
	log "github.com/Sirupsen/logrus"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	log.Info("Startng UDP client")

	udpAddr, err := net.ResolveUDPAddr("udp", ":8080")
	handleError(err)

	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	handleError(err)

	count := 1
	var buffer [512]byte
	for {
		_, err := udpConn.Write([]byte(strconv.Itoa(count)))
		handleError(err)
		time.Sleep(1 * time.Second)
		len, _, err1 := udpConn.ReadFromUDP(buffer[0:])
		handleError(err1)
		log.Info(string(buffer[0:len]))

		count = count + 1
	}
}

func handleError(err error) {
	if err != nil {
		log.Error("Fatal error: ", err.Error())
		os.Exit(1)
	}
}
