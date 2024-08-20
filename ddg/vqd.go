package ddg

import (
	"bytes"
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

func GetSearchVQD(query string) (string, error) {
	u, _ := url.Parse(BASE)

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

func getChatHeaderWithVQD() (http.Header, error) {
	req, _ := http.NewRequest(http.MethodGet, CHAT_STATUS, nil)
	req.Header.Set("x-vqd-accept", "1")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return nil, errors.New(string(body))
	}

	//vqd := res.Header.Get("x-vqd-4")

	//if vqd == "" {
	//	return nil, errors.New("no VQD found")
	//}

	return res.Header, nil
}
