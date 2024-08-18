package ddg

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
)

func TextSearch(query string) (TextResults, error) {
	vqd, err := getSearchVQD(query)

	if err != nil {
		return TextResults{}, err
	}

	log.Println("Search VQD:", vqd)

	u, _ := url.Parse(TEXT_ENDPOINT)

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

	req.Header.Set("Host", u.Host)

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

func ImageSearch(query string) (ImageResults, error) {
	vqd, err := getSearchVQD(query)

	if err != nil {
		return ImageResults{}, err
	}

	log.Println("Search VQD:", vqd)

	u, _ := url.Parse(IMAGE_ENDPOINT)

	q := u.Query()

	q.Set("q", query) // Query
	q.Set("vqd", vqd)
	q.Set("o", "json")         // OUTPUT: json, html
	q.Set("l", "us-en")        // REGION: wt-wt, us-en, uk-en, ru-ru
	q.Set("p", SAFE_SEARCH_ON) // SAFE SEARCH: on, moderate off
	q.Set("s", "0")
	q.Set("f", ",,,,,")

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)

	//req.Header.Set("Host", "links.duckduckgo.com")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return ImageResults{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return ImageResults{}, errors.New(res.Status)
		}
		return ImageResults{}, errors.New(string(resBody))
	}

	var body ImageResults

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return ImageResults{}, errors.New("Unable to parse response into ImageSearch")
	}

	return body, nil
}

func Chat(content string) (io.ReadCloser, error) {

	vqd, err := getChatVQD()

	log.Println("VQD found for Chat", vqd)

	if err != nil {
		return nil, err
	}

	u, _ := url.Parse(CHAT_ENDPOINT)

	prompt := Prompt{
		Model: LLAMA_3_70B,
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
