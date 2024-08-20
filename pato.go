package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/seanvelasco/pato/ddg"
	"log"
	"net/url"
	"strings"
)

func generateCompleteAnswer(prompt string) (string, error) {
	res, err := ddg.Chat(prompt)

	if err != nil {
		return "", err
	}

	defer res.Close()

	scanner := bufio.NewScanner(res)

	var wholeResponse string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "data:") {
			data := strings.TrimPrefix(line, "data:")
			if strings.HasSuffix(data, "[DONE]") {
				break
			}
			var body MessageSSE
			if err := json.Unmarshal([]byte(data), &body); err != nil {
				log.Println("Unable to unmarshal SSE", err)
			}
			wholeResponse += body.Message
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return wholeResponse, nil
}

func processQuery(query string) (string, []string, error) {
	log.Println("Getting VQD")
	token, err := ddg.GetSearchVQD(query)

	if err != nil {
		return "", nil, err
	}

	log.Println("Retrieved VQD")

	resultsChan := make(chan ddg.TextResults)
	imagesChan := make(chan ddg.ImageResults)
	errChan := make(chan error)

	go func() {
		log.Println("Starting to text search")
		results, err := ddg.SearchText(query, token)
		if err != nil {
			errChan <- err
			return
		}
		log.Println("Retrieved results")
		resultsChan <- results
	}()

	go func() {
		log.Println("Starting to image search")
		images, err := ddg.SearchImages(query, token)
		if err != nil {
			errChan <- err
			return
		}
		log.Println("Retrieved images")
		imagesChan <- images
	}()

	var imageURLs []string
	var reference string

	select {
	case results := <-resultsChan:
		for i, r := range results.Results {
			//results.Results[i].Title = removeHTMLBTag(r.Title)
			//results.Results[i].Body = formatString(r.Body)
			if r.Body != "" {
				results.Results[i].Body = formatString(r.Body)
				reference += fmt.Sprintf("%s: %s (%s)\n\n", removeHTMLBTag(r.Title), formatString(r.Body), r.URL)
			}
		}
	case images := <-imagesChan:
		for _, image := range images.Results {
			validURL, _ := url.Parse(image.URL)
			if validURL != nil {
				imageURLs = append(imageURLs, validURL.String())
			}
		}
	case err := <-errChan:
		return "", nil, err
	}

	prompt := fmt.Sprintf(
		"Answer the question or find more information to %s. "+
			"The format of each source is 'TITLE: CONTENT (SOURCE)'. "+
			"The resulting answer should be an answer to the query, indicated with indices [n], "+
			"and the URLs of the references used should be listed (1., 2., 3., ..., n.) as the last part. "+
			"Strictly do not use a source if it has no relation to the main answer or if it's off-topic. "+
			"Strictly do not include references if they haven't been used as indices on the answer. "+
			"Do not return the query as the heading of the answer. "+
			"Do not attempt to interpret the question. "+
			"If the query is one word, then it's just one word. Do not make it longer."+
			"Supplement the answer with %s. This should not be the the main content of the answer. The answer must be factual and positive. Include as little news as possible. You may introduce a fun fact. ",
		query, reference)

	answer, err := generateCompleteAnswer(prompt)

	if err != nil {
		return "", nil, err
	}

	return answer, imageURLs, nil
}
