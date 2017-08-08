package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	// read text and convert the type as io.Reader interface
	reader := strings.NewReader("ʕ ◔ϖ◔ʔ < Sample Text")
	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		// fail to transmit
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
