package ddg

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func extractVQD(body []byte) (string, error) {

	re := regexp.MustCompile(`vqd=([^,]+)`)

	matches := re.FindSubmatch(body)

	if len(matches) < 2 {
		return "", errors.New("no VQD found")
	}

	return strings.Trim(string(matches[1]), "\""), nil
}

func getVQD(query string) (string, error) {
	u, _ := url.Parse("https://duckduckgo.com")

	body := &bytes.Buffer{} // new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	writer.WriteField("q", query)

	err := writer.Close()

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), body)

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)

	return extractVQD(bodyBytes)
}

func TextSearch(query string) (TextResults, error) {
	vqd, err := getVQD(query)

	if err != nil {
		return TextResults{}, err
	}

	u, _ := url.Parse("https://links.duckduckgo.com/d.js")

	q := u.Query()

	q.Set("q", query)
	q.Set("o", "json")
	q.Set("vqd", vqd)
	q.Set("kl", "us-en")
	q.Set("l", "us-en")
	// Unsure if these params are needed
	q.Set("p", "")
	q.Set("s", "0")
	q.Set("df", "d")
	q.Set("bing_market", "EN-US")

	u.RawQuery = q.Encode()

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)

	req.Header.Set("Host", "links.duckduckgo.com")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return TextResults{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		resBody, _ := io.ReadAll(res.Body)
		return TextResults{}, errors.New(string(resBody))
	}

	// DDG can also return an error while having a http.StatusOK status code

	var body TextResults

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return TextResults{}, errors.New("Unable to parse response into TextSearch")
	}

	return body, nil
}

func Chat(content string) (io.ReadCloser, error) {

	vqd, err := getVQD(content)

	if err != nil {
		return nil, err
	}

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
