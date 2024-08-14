package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func vqd(prompt Prompt) (string, error) {
	u, _ := url.Parse("https://duckduckgo.com/")
	q := u.Query()
	q.Set("q", "DuckDuckGo AI Chat")
	q.Set("ia", "chat")
	q.Set("duckai", "1")
	q.Set("atb", "v425-1")
	u.RawQuery = q.Encode()

	return "4-264950795912270085190654548782964737427", nil
}

func ai(content string) (io.ReadCloser, error) {
	u, _ := url.Parse("https://duckduckgo.com/duckchat/v1/chat")

	prompt := Prompt{
		Model: "meta-llama/Meta-Llama-3.1-70B-Instruct-Turbo",
		Messages: []PromptMessage{
			{
				Role:    "user",
				Content: content,
			},
		},
	}

	body, _ := json.Marshal(prompt)

	//vqd, err := vqd(prompt)
	//
	//if err != nil {
	//	return nil, err
	//}

	req, _ := http.NewRequest("POST", u.String(), bytes.NewReader(body))

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:129.0) Gecko/20100101 Firefox/129.0")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-vqd-4", "4-267030941856431541085535292874624219683")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Priority", "u=4")
	req.Header.Set("credentials", "include")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", res.StatusCode)
	}

	return res.Body, nil
}
