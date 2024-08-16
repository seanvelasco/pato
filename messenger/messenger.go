package messenger

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func SendMessage(pageID string, recipientID string, messageID string, text string) (SendMessageResponse, error) {
	u, err := url.Parse(fmt.Sprintf("https://graph.facebook.com/v20.0/%s/messages", pageID))
	if err != nil {
		return SendMessageResponse{}, err
	}

	q := u.Query()
	q.Set("access_token", os.Getenv("META_PAGE_ACCESS_TOKEN"))
	u.RawQuery = q.Encode()

	response := SendMessageRequest{
		Recipient: User{
			ID: recipientID,
		},
		MessagingType: "RESPONSE",
		Message: Message{
			Text: text,
			ReplyTo: MessageReplyTo{
				MID: messageID,
			},
		},
	}

	reqBody, err := json.Marshal(response)

	if err != nil {
		return SendMessageResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(reqBody))

	if err != nil {
		return SendMessageResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}

	res, err := client.Do(req)

	if err != nil {
		return SendMessageResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		resBody, _ := io.ReadAll(res.Body)
		return SendMessageResponse{}, errors.New(string(resBody))
	}

	var resBody SendMessageResponse

	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		return SendMessageResponse{}, err
	}

	return resBody, nil
}
