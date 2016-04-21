package main

import (
	log "github.com/Sirupsen/logrus"
	"net"
	"time"
)

func main() {
	log.Info("Starting TCP client")
	for {
		conn, err1 := net.Dial("tcp4", ":8080")
		if err1 != nil {
			log.Error("Fatal Error: ", err1.Error())
		}
		var buffer [512]byte
		_, err2 := conn.Read(buffer[0:])
		if err2 != nil {
			log.Error("Fatal Error: ", err2.Error())
		}
		log.Info(string(buffer[0:]))
		time.Sleep(1 * time.Second)
	}
}
