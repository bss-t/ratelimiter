package main

import (
	"log"
	"net/http"

	"github.com/bss-t/ratelimiter/pkg/endpoints"
)

func main() {
	http.HandleFunc("/consume", endpoints.HandleConsume) // Endpoint to consume a token
	http.HandleFunc("/status", endpoints.HandleStatus)   // Endpoint to check bucket status

	log.Println("Token Bucket Service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
