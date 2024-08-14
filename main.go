package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlePing)
	mux.HandleFunc("GET /messenger", handleVerification)
	mux.HandleFunc("POST /messenger", handleMessages)

	http.ListenAndServe(":8080", mux)
}
