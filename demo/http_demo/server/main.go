package main

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

func main() {

	//router
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", Hello)

	//server
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		panic(err.Error())
	}

}

func Hello(resp http.ResponseWriter, req *http.Request) {
	_, err := resp.Write([]byte("Hello:" + req.RemoteAddr))
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", errors.New(err.Error())))
	}
}
