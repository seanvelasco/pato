package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/seanvelasco/pato/messenger"
	"github.com/seanvelasco/pato/telegram"
	"log"
	"net/http"
	"os"
	"strconv"
)

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

	var body messenger.MessageEvent

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	defer r.Body.Close()

	for _, entry := range body.Entry {
		if entry.Messaging != nil {
			for _, m := range entry.Messaging {
				go func() {
					answer, _, err := processQuery(m.Message.Text)
					if err != nil {
						log.Println("Unable to generate completion:", err)
						if _, err := messenger.SendMessage(m.Recipient.ID, m.Sender.ID, m.Message.MID, BREAK); err != nil {
							log.Println("Unable to send a Messenger message:", err)
						}
						return
					}
					if _, err := messenger.SendMessage(m.Recipient.ID, m.Sender.ID, m.Message.MID, answer); err != nil {
						log.Println("Unable to send a Messenger message:", err)
					}
				}()
			}
		}
	}
}

func handleTelegramMessages(w http.ResponseWriter, r *http.Request) {
	var body telegram.Update

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	go func() {

		chatID := strconv.Itoa(body.Message.Chat.ID)
		messageID := strconv.Itoa(body.Message.MessageID)

		answer, _, err := processQuery(body.Message.Text)

		if err != nil {
			log.Println("Unable to generate completion:", err)
			if _, err := telegram.SendMessage(chatID, BREAK, messageID); err != nil {
				log.Println("Unable to send a Telegram message:", err)
			}
			return
		}
		if _, err := telegram.SendMessage(chatID, answer, messageID); err != nil {
			log.Println("Unable to send a Telegram message:", err)
		}
	}()
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
