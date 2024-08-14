package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

//res, err := ai(body.Message.Text)
//
//if err != nil {
//log.Fatal(err)
//}
//
//defer res.Close()

//scanner := bufio.NewScanner(res)
//
//var wholeResponse string
//
//for scanner.Scan() {
// line := scanner.Text()
//if strings.HasPrefix(line, "data:") {
//data := strings.TrimPrefix(line, "data:")
//if strings.HasSuffix(data, "[DONE]") {
//break
//}
//var body MessageSSE
//if err := json.Unmarshal([]byte(data), &body); err != nil {
//log.Fatal("Unable to unmarshal SSE", err)
//}
//wholeResponse += body.Message
//}
//}
//
//if err := scanner.Err(); err != nil {
//log.Fatal(err)
//}

func handlePing(w http.ResponseWriter, r *http.Request) {

}

func validateSignature(signature string, secret string) bool {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(signature))
	expectedMAC := hex.EncodeToString(h.Sum(nil))
	return hmac.Equal([]byte(expectedMAC), []byte(secret))
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	//signature := r.Header.Get("x-hub-signature-256")
	//appSecret := os.Getenv("META_APP_SECRET")
	//
	//if signature == "" || appSecret == "" {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//
	//if signature != appSecret {
	//	w.WriteHeader(http.StatusForbidden)
	//	return
	//}
	//
	//signatureUnicode := strings.ReplaceAll(signature, "ä", "\\u00e4")
	//signaturePart := strings.Split(signatureUnicode, ":")[1]
	//
	//if validateSignature(signaturePart, appSecret) != true {
	//	w.WriteHeader(http.StatusForbidden)
	//	return
	//}

	var body WebhookEvent

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("EVENT RECEIVED"))

	message := body.Value.Message.Text

	log.Println(message)

	//for _, entry := range body. {
	//	for _, messaging := range entry.Messaging {
	//		log.Println(messaging.Message)
	//	}
	//}

}

func handleMessagingPostbacks(w http.ResponseWriter, r *http.Request) {

}

func handleVerification(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mode := q.Get("hub.mode")
	token := q.Get("hub.verify_token")
	challenge := q.Get("hub.challenge")

	if mode != "" && token != "" {
		if mode == "subscribe" && token == os.Getenv("META_MESSENGER_VERIFY_TOKEN") {
			w.Write([]byte(challenge))
			return
		}
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
