package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func handleMessages(w http.ResponseWriter, r *http.Request) {

	var body SendMessageRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	res, err := ai(body.Message.Text)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	scanner := bufio.NewScanner(res)

	var wholeResponse string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "data:") {
			data := strings.TrimPrefix(line, "data:")
			if strings.HasSuffix(data, "[DONE]") {
				break
			}
			var body MessageSSE
			if err := json.Unmarshal([]byte(data), &body); err != nil {
				log.Fatal("Unable to unmarshal SSE", err)
			}
			wholeResponse += body.Message
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(wholeResponse))
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

	mux := http.NewServeMux()

	mux.HandleFunc("POST /messages", handleMessages)
	mux.HandleFunc("/verify", handleVerification)

	http.ListenAndServe(":8080", mux)
}
