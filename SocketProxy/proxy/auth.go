package proxy

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

const socks5Ver = 0x05

func auth(reader *bufio.Reader, client net.Conn) (err error) {
	// the format of auth response is version + method size + method
	ver, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read version fail %v", err)
	}
	if ver != socks5Ver {
		return fmt.Errorf("version not support %v", err)
	}
	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read method size fail %v", err)
	}
	method := make([]byte, methodSize)
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method fail %v", err)
	}
	log.Println("version: ", ver, "method", method)

	// send response to confirm
	_, err = client.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("send response fail %v", err)
	}
	return nil
}
