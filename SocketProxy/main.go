package SocketProxy

import (
	"Go_Project/SocketProxy/proxy"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "localhost:57028")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Println("Acceptation Failure %v", err)
			continue
		}
		go proxy.Process(client)

	}
}
