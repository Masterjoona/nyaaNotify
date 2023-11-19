package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ShortenedURL struct {
	Url string `json:"url"`
}

func FetchNyaa(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	return string(content)

}

func ShortenURL(url string, shortenerToken string, shortenerURL string) (string, error) {
	payload := []byte(`{"url":"` + url + `"}`)
	req, err := http.NewRequest("POST", shortenerURL, bytes.NewBuffer(payload))
	if err != nil {
		Logger("Error creating request: " + err.Error())
		return "", err
	}

	req.Header.Set("Authorization", shortenerToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Max-Views", "5")
	// req.Header.Set("No-JSON", "true")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		Logger("Error sending request: " + err.Error())
		return "", err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		Logger("Error reading response: " + err.Error())
		return "", err
	}

	shortenedURL := ShortenedURL{}
	err = json.Unmarshal(content, &shortenedURL)
	if err != nil {
		Logger("Error unmarshalling response: " + err.Error())
		return "", err
	}
	// return string(content), nil
	return shortenedURL.Url, nil
}
