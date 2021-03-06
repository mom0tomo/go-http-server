package main

import (
	//"io/ioutil"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// use like "$ curl -G --data-urlencode "query=hello world" http:/localhost:18888"
func main() {
	// declare query string
	values := url.Values{
		// make query string
		"query": {"hello world"},
	}

	// add query string at the end of URL
	resp, err := http.Get("http://localhost:18888" + "?" + values.Encode())
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	// read body content with bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}
