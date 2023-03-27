package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}
