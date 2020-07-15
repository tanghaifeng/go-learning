package http_01

import (
	"fmt"
	"net/http"
)

func HttpServer() {
	var err error
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		s := fmt.Sprintf("hello http")
		fmt.Fprintln(writer, s)
	})
	if err = http.ListenAndServe(":8080", nil); err != nil {
		panic("ListenAndServe err")
	}

}
