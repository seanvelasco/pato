package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func SendMessage(chatID string, text string, replyToMessageID string) (Message, error) {
	u, err := url.Parse(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", os.Getenv("TELEGRAM_BOT_TOKEN")))
	if err != nil {
		return Message{}, err
	}

	q := u.Query()
	q.Set("chat_id", chatID)
	q.Set("text", text)
	q.Set("reply_to_message_id", replyToMessageID)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)

	if err != nil {
		return Message{}, err
	}

	client := &http.Client{Timeout: 5 * time.Second}

	res, err := client.Do(req)

	if err != nil {
		return Message{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		resBody, _ := io.ReadAll(res.Body)
		return Message{}, errors.New(string(resBody))
	}

	var sentMessage Message

	if err := json.NewDecoder(res.Body).Decode(&sentMessage); err != nil {
		return Message{}, err
	}

	return sentMessage, nil
}
