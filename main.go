package main

import "net/http"

func handleMessages(w http.ResponseWriter, r *http.Request) {

}

func handleMessagingPostbacks(w http.ResponseWriter, r *http.Request) {

}

func handleVerification(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mode := q.Get("hub.mode")
	token := q.Get("hub.verify_token")
	challenge := q.Get("hub.challenge")

	if mode != "" && token != "" {
		if mode == "subscribe" && token == "123abcxyz" {
			_, _ = w.Write([]byte(challenge))
			return
		}
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func main() {
	ai()
	mux := http.NewServeMux()

	mux.HandleFunc("/messages", handleMessages)
	mux.HandleFunc("/verify", handleVerification)

	http.ListenAndServe(":8080", mux)
}
