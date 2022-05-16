package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const Header string = "12345678"

func Encode(conn io.Writer, content []byte) error {

	err := binary.Write(conn, binary.BigEndian, []byte(Header))
	if err != nil {
		return err
	}

	length := uint32(len(content))
	if err = binary.Write(conn, binary.BigEndian, length); err != nil {
		return err
	}

	return binary.Write(conn, binary.BigEndian, content)
}

func main() {

	n := 2

	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer conn.Close()
	for i := 0; i <= n; i++ {
		content := []byte(fmt.Sprintf("Hello World %d\n", i))

		if err = Encode(conn, content); err != nil {
			fmt.Println(err.Error())
		}
	}
}
