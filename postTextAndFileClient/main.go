package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	// declare buffer to store byte line
	var buffer bytes.Buffer

	// make writer
	writer := multipart.NewWriter(&buffer)
	// register non file fields
	writer.WriteField("name", "John Lennon")

	// make io.Writer to read each files
	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}

	// open files
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		// fail to open files
		panic(err)
	}

	defer readFile.Close()
	// copy all the contents of files
	io.Copy(fileWriter, readFile)
	// close multi part(io.Writer) and write all the contents in the buffer
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
