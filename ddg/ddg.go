package ddg

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func getVQD(prompt Prompt) (string, error) {
	u, _ := url.Parse("https://duckduckgo.com/")
	q := u.Query()
	q.Set("q", "DuckDuckGo AI Chat")
	q.Set("ia", "chat")
	q.Set("duckai", "1")
	q.Set("atb", "v425-1")
	u.RawQuery = q.Encode()

	return "128763328289690301316563025964306385877", nil
}

func AI(content string) (io.ReadCloser, error) {
	u, _ := url.Parse("https://duckduckgo.com/duckchat/v1/chat")

	prompt := Prompt{
		Model: "meta-llama/Meta-Llama-3.1-70B-Instruct-Turbo",
		Messages: []Message{
			{
				Role:    "user",
				Content: content,
			},
		},
	}

	body, _ := json.Marshal(prompt)

	vqd, err := getVQD(prompt)

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", u.String(), bytes.NewReader(body))

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:129.0) Gecko/20100101 Firefox/129.0")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-vqd-4", vqd)
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
		resBody, _ := io.ReadAll(res.Body)
		return nil, errors.New(string(resBody))
	}

	return res.Body, nil
}
