package telegram

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
)

func SendMessage(chatID string, text string) (Message, error) {
	u, _ := url.Parse("https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN"))
	q := u.Query()
	q.Set("chat_id", os.Getenv(chatID))
	q.Set("text", text)
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())

	if err != nil {
		return Message{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Message{}, errors.New(res.Status)
	}

	var sentMessage Message

	if err := json.NewDecoder(res.Body).Decode(&sentMessage); err != nil {
		log.Println(err)
	}

	return sentMessage, nil
}
