package messenger

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
)

func SendMessage(pageID string, recipientID string, text string) (SendMessageResponse, error) {
	u, err := url.Parse("https://graph.facebook.com/v20.0/" + pageID + "/messages")
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
		Message: struct {
			Text string `json:"text"`
		}{
			Text: text,
		},
	}

	reqBody, _ := json.Marshal(response)

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(reqBody))

	if err != nil {
		return SendMessageResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return SendMessageResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return SendMessageResponse{}, errors.New("Bad status code: " + res.Status)
	}

	var resBody SendMessageResponse

	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		return SendMessageResponse{}, err
	}

	return resBody, nil
}
