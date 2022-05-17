package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	input := bufio.NewReader(os.Stdin)

	str, _ := input.ReadString('\n')
	str = strings.TrimSpace(str) //trimspace可以去除头尾两端的全部\n,\r,\t等等空白rune
	fmt.Println(str + "hello")
}
