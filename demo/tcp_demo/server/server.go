package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

const Header = "12345678"

func Decode(conn io.Reader) ([]byte, error) {
	var (
		err       error
		HeaderBuf []byte
		lengthBuf []byte
		content   []byte
		length    uint32
	)
	HeaderBuf = make([]byte, len(Header))
	if _, err = io.ReadFull(conn, HeaderBuf); err != nil {
		return nil, err
	}

	if Header != string(HeaderBuf) {
		return nil, errors.New("Header Error")
	}

	lengthBuf = make([]byte, binary.Size(length))
	if _, err = io.ReadFull(conn, lengthBuf); err != nil {
		return nil, err
	}

	length = binary.BigEndian.Uint32(lengthBuf)

	content = make([]byte, length)
	if _, err = io.ReadFull(conn, content); err != nil {
		return nil, err
	}

	return content, nil

}

func main() {

	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		} else {
			go Process(conn)
		}

	}
}

func Process(conn net.Conn) {
	defer conn.Close()
	for {
		content, err := Decode(conn)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(content))
	}
}
