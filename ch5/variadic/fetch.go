package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		response.Body.Close()
		fmt.Printf("%s", body)
	}
}
