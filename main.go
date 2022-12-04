package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/t1", func(writer http.ResponseWriter, request *http.Request) {
		//fmt.Println(request.Header)
		for k, v := range request.Header {
			fmt.Println(k, v)
		}
		writer.WriteHeader(http.StatusOK)
		_, err := writer.Write([]byte("ttttttttttttttttttttttt1"))
		if err != nil {
			fmt.Println("write err", err)
			return
		}
	})
	if err := http.ListenAndServe(":11111", &ss{}); err != nil {
		fmt.Println("start error", err)
	}

}

type ss struct {
}

func (s *ss) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Println("in /")
	default:
		fmt.Println("other")
	}
}
