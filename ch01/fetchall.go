package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetch(url string, channel chan string) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		channel <- fmt.Sprint(err) // send error to channel
		return
	}
	elapsed := time.Since(start)
	if nbytes, err := io.Copy(io.Discard, response.Body); err != nil {
		channel <- fmt.Sprintf("Error while reading %s: %v", url, err)
	} else {
		channel <- fmt.Sprintf("%.2fs %7d %s", elapsed.Seconds(), nbytes, url)
	}
	response.Body.Close()
}

func main() {
	urls := os.Args[1:]
	channel := make(chan string)
	for _, url := range urls {
		go fetch(url, channel)
	}
	for range urls {
		fmt.Println(<-channel)
	}
}
