package proxy

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const cmdBind = 0x01
const atypeIPV4 = 0x01
const atypeHost = 0x03
const atypeIPV6 = 0x04

func connect(reader *bufio.Reader, client net.Conn) (err error) {
	// connect header = Ver + CMD + RSV + ATYP + DST.ADDR + DST.PORT
	// CMD: connect request present by 0x01
	// RSV: for extra info
	// ATYP: address type
	// 		if == 0x01 then is a IPV4
	//		if == 0x03 then is a domain
	// DST.ADDR: the address
	// DST.Port: the port
	buf := make([]byte, 4) // we need know the ATYP to judge what address is
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return fmt.Errorf("read header fail %v", err)
	}
	ver, cmd, atyp := buf[0], buf[1], buf[3]
	if ver != socks5Ver {
		return fmt.Errorf("version not support %v", err)
	}
	if cmd != cmdBind {
		return fmt.Errorf("cmd not support %v", err)
	}
	var addr string
	switch atyp {
	case atypeIPV4:
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return fmt.Errorf("read address fail %v", err)
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])

	case atypeHost:
		hostSize, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("read host size fail %v", err)
		}
		host := make([]byte, hostSize)
		_, err = io.ReadFull(reader, host)
		if err != nil {
			return fmt.Errorf("read host fail %v", err)
		}
		addr = string(host)

	case atypeIPV6:
		return errors.New("IPV6: Not support yet")
		// not support yet
		//_, err = io.ReadFull(reader, buf)
		//if err != nil {
		//	return fmt.Errorf("read address fail %v", err)
		//}
		//addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	default:
		return errors.New("invalid ATYP")
	}
	_, err = io.ReadFull(reader, buf[:2]) // since port need 2 bytes
	if err != nil {
		return fmt.Errorf("read port fail %v", err)
	}
	port := binary.BigEndian.Uint16(buf[:2])

	// relay
	dest, err := net.Dial("tcp", fmt.Sprintf("%v:%v", addr, port))
	if err != nil {
		return fmt.Errorf("dial destation failed: %v", err)
	}
	defer Cclose(dest)
	log.Println("dial", addr, ":", port)

	// need to send back response, ver + REP + RSV + ATYPE + BND.ADDR(4 bytes) + BND.PORT(2 bytes)
	// here we do proxy locally
	_, err = client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
	if err != nil {
		return fmt.Errorf("write response fail %v", err)
	}
	// use copy function to do data exchange
	// but connect function will return immediately so that we need use context module to control

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		_, _ = io.Copy(dest, reader)
		cancel()
	}()
	go func() {
		_, _ = io.Copy(client, dest)
		cancel()
	}()
	<-ctx.Done()
	return nil
}
