package main

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		fmt.Printf("%+v\n", errors.Wrap(err, "listen failed"))
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("%+v\n", errors.Wrap(err, "conn failed"))
			continue
		}

		go Process(conn)

	}
}

func Process(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 128)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("%+v\n", errors.Wrap(err, "read from conn error:"))
			break
		}

		fmt.Println(string(buf[:n]))
	}
}
