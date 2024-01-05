package main

import (
	"io"
	"net/http"
)

func FetchNyaa(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		Logger("Error fetching nyaa: " + err.Error())
		panic(err)
	}

	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)

	if err != nil {
		Logger("Error reading nyaa response: " + err.Error())
		panic(err)
	}
	return string(content)

}
