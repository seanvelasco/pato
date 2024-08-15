package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlePing)
	mux.HandleFunc("GET /messenger", handleVerification)
	mux.HandleFunc("POST /messenger", handleMessages)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}
