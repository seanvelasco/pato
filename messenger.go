package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

func send_message(pageID string, recipientID string, text string) (SendMessageResponse, error) {
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

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return SendMessageResponse{}, err
	}

	defer res.Body.Close()

	var res_body SendMessageResponse

	if err := json.NewDecoder(res.Body).Decode(&res_body); err != nil {
		return SendMessageResponse{}, err
	}

	return res_body, nil
}
