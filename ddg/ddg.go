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

func SearchAssist(query string, vqd string) (SearchAssistResults, error) {

	u, _ := url.Parse(ASSIST_ENDPOINT)
	q := u.Query()
	q.Set("q", query)
	q.Set("vqd", vqd)
	q.Set("shadow_webrag", "1")
	u.RawQuery = q.Encode()

	body, _ := json.Marshal(struct {
		q string `json:"q"`
	}{
		q: query,
	})

	req, err := http.NewRequest("GET", u.String(), bytes.NewReader(body))

	if err != nil {
		return SearchAssistResults{}, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return SearchAssistResults{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		resBody, _ := io.ReadAll(res.Body)
		return SearchAssistResults{}, errors.New(string(resBody))
	}

	var results SearchAssistResults

	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		resBody, _ := io.ReadAll(res.Body)
		log.Println(string(resBody), res.StatusCode)
		return SearchAssistResults{}, errors.New("unable to parse response into SearchAssistResults")
	}

	return results, nil
}

func SearchSuggestions(query string) (SuggestionResults, error) {
	u, _ := url.Parse(SUGGESTIONS_ENDPOINT)
	q := u.Query()
	q.Set("q", query)
	q.Set("kl", "en-us")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var body SuggestionResults

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return SuggestionResults{}, errors.New("unable to parse response into SuggestionResults")
	}

	return body, nil
}

func SearchText(query string, vqd string) (TextResults, error) {

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

func SearchImages(query string, vqd string) (ImageResults, error) {

	u, _ := url.Parse(IMAGE_ENDPOINT)

	q := u.Query()

	q.Set("q", query) // Query
	q.Set("vqd", vqd)
	q.Set("o", "json")         // OUTPUT: json, html
	q.Set("l", "us-en")        // REGION: wt-wt, us-en, uk-en, ru-ru
	q.Set("p", SAFE_SEARCH_ON) // SAFE SEARCH: on, moderate off
	q.Set("s", "0")
	q.Set("f", ",,,,,")

	u.RawQuery = q.Encode()

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)

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

	vqd, err := getChatHeaderWithVQD()

	if err != nil {
		return nil, err
	}

	log.Println("VQD found for Chat", vqd.Get("x-vqd-4"))

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

	//req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Host", u.Host)
	//req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:129.0) Gecko/20100101 Firefox/129.0")
	//req.Header.Set("x-vqd-4", "4-257864896927281724418427422564929565235")
	//req.Header.Set("Accept", "text/event-stream")
	req.Header = vqd

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
