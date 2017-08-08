package main

import (
	"log"
	"net/http"
)

// use like "$ curl --head http://localhost:18888"
func main() {
	// get header content
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	// print status code
	log.Println("Status:", resp.Status)
	// print header content
	log.Println("Headers:", resp.Header)
}
