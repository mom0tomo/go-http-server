package main

import (
	"log"
	"net/http"
	"os"
)

// use like "$ curl -T main.go -H "Content-type: text/plain" http://localhost:18888"
func main() {
	// make os.File object as io.Reader interface
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		// fail to transmit
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
