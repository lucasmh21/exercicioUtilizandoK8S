package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,greeting("Code.education Rocks!"), request.URL.Path[1:])
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func greeting(texto string) string {
	return "<b>" + texto + "</b>"
}
