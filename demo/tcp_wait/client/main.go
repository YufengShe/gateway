package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Printf("%+v\n", errors.Wrap(err, "conn failed"))
		return
	}
	defer conn.Close()

	//input
	input := bufio.NewReader(os.Stdin)

	for {
		str, err := input.ReadString('\n')
		if err != nil {
			fmt.Printf("%+v\n", errors.Wrap(err, "input error"))
			continue
		}

		strimStr := strings.TrimSpace(str)
		if strimStr == "Q" {
			break
		}

		_, err = conn.Write([]byte(strimStr))
		if err != nil {
			fmt.Printf("%+v\n", errors.Wrap(err, "write error"))
			break
		}
	}

}
