package main

import (
	"log"
	"net/http"
)

func main() {

	zzz, err := generateAnswer("are you an AI?")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(zzz)

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlePing)
	mux.HandleFunc("GET /messenger", handleVerification)
	mux.HandleFunc("POST /messenger", handleMessages)

	http.ListenAndServe(":8080", mux)
}
