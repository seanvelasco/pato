package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Prompt struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func ai() {
	u, _ := url.Parse("https://duckduckgo.com/duckchat/v1/chat")

	body, err := json.Marshal(Prompt{
		Model: "gpt-4o-mini",
		Messages: []Message{
			{
				Role:    "user",
				Content: "are you an ai",
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("POST", u.String(), bytes.NewReader(body))

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:129.0) Gecko/20100101 Firefox/129.0")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-vqd-4", "4-264950795912270085190654548782964737427")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
