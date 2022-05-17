package main

import (
	E "errors"
	"fmt"
	"github.com/pkg/errors"
)

func main() {

	err := E.New("123")
	err = errors.Wrap(err, "test")
	fmt.Println(fmt.Sprintf("%+v", err))
}
