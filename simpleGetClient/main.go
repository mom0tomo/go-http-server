package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// bodyの内容をstring型で取得
	log.Println(string(body))
	// ステータスコードをstring型で取得
	log.Println("Status:", resp.Status)
	// ヘッダーをmap型で取得
	log.Println("Headers:", resp.Header)
	// Content-Typeの最初の要素をstring型で取得
	log.Println("Content-Type:", resp.Header.Get("Content-Type"))
}
