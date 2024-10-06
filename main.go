package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"syscall"
)

// main function does something
func main() {
	Response, err := http.Get("https://jsonplaceholder.typicode.com/posts") // Variable name should be camelCase
	if err != nil {
		log.Println("Error fetching URL:", err)
		return
	}

	file, err := os.OpenFile("test.txt", syscall.O_RDWR|syscall.O_CREAT|syscall.O_TRUNC, 0666)

	if err != nil {
		log.Println("Error opening file:", err)
	}

	defer file.Close()
	defer Response.Body.Close()

	if Response.StatusCode != http.StatusOK {
		log.Println("Unexpected status code:", Response.StatusCode)
		return
	}

	body, err := io.ReadAll(Response.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	file.Write(body)
}
