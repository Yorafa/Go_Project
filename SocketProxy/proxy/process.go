package proxy

import (
	"bufio"
	"log"
	"net"
)

func Cclose(client net.Conn) {
	err := client.Close()
	if err != nil {
		log.Println("close fail")
	}
}

func Process(client net.Conn) {
	defer Cclose(client)
	reader := bufio.NewReader(client)
	err := auth(reader, client)
	if err != nil {
		log.Printf("client %v auth failed: %v", client.RemoteAddr(), err)
		return
	}
	log.Println("Authorization Success")
	err = connect(reader, client)
	if err != nil {
		log.Printf("client %v connect failed: %v", client.RemoteAddr(), err)
		return
	}
}

func Repeater(client net.Conn) {
	defer Cclose(client)
	reader := bufio.NewReader(client)
	for {
		buf, err := reader.ReadByte()
		if err != nil {
			break
		}
		_, err = client.Write([]byte{buf})
		if err != nil {
			break
		}
	}
}
