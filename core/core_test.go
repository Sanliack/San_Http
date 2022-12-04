package core

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_core(t *testing.T) {
	s := NewEngine()
	s.Get("/aa", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("555555555555555555"))
	})

	if e := s.Server(":11111"); e != nil {
		fmt.Println(e)
	}
}
