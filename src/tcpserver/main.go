package main

import (
	"fmt"
	"log"
	"tcpserver/server"
	"time"
)

func main() {
	tcpServer := server.NewServer(":8080")

	go func() {
		for msg := range tcpServer.Msgch {
			fmt.Println(string(msg))
		}
	}()

	go func() {
		<-time.After(1 * time.Minute)
		tcpServer.Close()
	}()

	log.Fatal(tcpServer.Serve())

}
